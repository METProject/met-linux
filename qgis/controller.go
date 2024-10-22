package qgis

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/imide/met-linux/util"
	"go.uber.org/zap"
)

// QGISController handles interactions with the Python QGIS script
type QGISController struct {
	scriptPath   string
	sem          chan struct{}
	outputFolder string
	config       *util.Config
}

// NewQGISController creates a new QGIS controller instance
func NewQGISController(cfg *util.Config) (*QGISController, error) {
	scriptPath := filepath.Join("scripts", "qgiscontroller.py")
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("qgiscontroller.py not found at %s", scriptPath)
	}

	outputFolder := filepath.Join(cfg.Path.ScriptsFolder, "image_exports")
	if err := os.MkdirAll(outputFolder, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %v", err)
	}

	return &QGISController{
		scriptPath:   scriptPath,
		sem:          make(chan struct{}, cfg.Gen.Threads),
		outputFolder: outputFolder,
		config:       cfg,
	}, nil
}

// FixGeometry runs the fix_geometry function from the Python script
func (q *QGISController) FixGeometry(projectPath, algorithm string, parameters map[string]interface{}) error {
	// Convert parameters to Python dict string format
	paramStr := "{"
	for k, v := range parameters {
		switch val := v.(type) {
		case string:
			paramStr += fmt.Sprintf("'%s': '%v',", k, val)
		default:
			paramStr += fmt.Sprintf("'%s': %v,", k, val)
		}
	}
	paramStr += "}"

	q.sem <- struct{}{}        // Acquire semaphore
	defer func() { <-q.sem }() // Release semaphore

	cmd := exec.Command("python3", q.scriptPath, "fix_geometry",
		"--project-path", projectPath,
		"--algorithm", algorithm,
		"--parameters", paramStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		zap.L().Error("Failed to execute QGIS fix_geometry",
			zap.String("output", string(output)),
			zap.Error(err))
		return fmt.Errorf("QGIS fix_geometry failed: %v", err)
	}

	return nil
}

// ExportImageTile exports a single image tile
func (q *QGISController) ExportImageTile(projectPath string, xMin, xMax, yMin, yMax int,
	layerOutputName string, layers []string) error {

	q.sem <- struct{}{}        // Acquire semaphore
	defer func() { <-q.sem }() // Release semaphore

	// Calculate tile name and create output directory
	tile := calculateTile(xMin, yMax)
	tileOutputPath := getTileOutputPath(q.outputFolder, tile, layerOutputName)

	// Skip if tile already exists
	if _, err := os.Stat(tileOutputPath); err == nil {
		zap.L().Info("Skipping existing tile", zap.String("tile", tile))
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(tileOutputPath), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create tile directory: %v", err)
	}

	// Convert layers slice to Python tuple string format
	layersStr := "("
	for _, layer := range layers {
		layersStr += fmt.Sprintf("'%s',", layer)
	}
	layersStr += ")"

	cmd := exec.Command("python3", q.scriptPath, "export_image_tile",
		"--project-path", projectPath,
		"--block-per-tile", fmt.Sprintf("%d", q.config.Gen.BlocksPerTile),
		"--degree-per-tile", fmt.Sprintf("%d", q.config.Gen.DegreePerTile),
		"--x-min", fmt.Sprintf("%d", xMin),
		"--x-max", fmt.Sprintf("%d", xMax),
		"--y-min", fmt.Sprintf("%d", yMin),
		"--y-max", fmt.Sprintf("%d", yMax),
		"--output-path", tileOutputPath,
		"--layer-output-name", layerOutputName,
		"--layers", layersStr)

	output, err := cmd.CombinedOutput()
	if err != nil {
		zap.L().Error("Failed to execute QGIS export_image_tile",
			zap.String("output", string(output)),
			zap.Error(err))
		return fmt.Errorf("QGIS export_image_tile failed: %v", err)
	}

	return nil
}

// ExportImage exports multiple image tiles in parallel
func (q *QGISController) ExportImage(projectPath string,
	xRangeMin, xRangeMax, yRangeMin, yRangeMax int,
	layerOutputName string, layers []string) error {

	var wg sync.WaitGroup
	errChan := make(chan error, 1) // Buffer for first error
	done := make(chan struct{})

	// Start goroutine to collect errors
	go func() {
		for err := range errChan {
			if err != nil {
				zap.L().Error("Error in export image tile", zap.Error(err))
				close(done) // Signal to stop processing
				return
			}
		}
	}()

	// Process tiles in parallel
	for xMin := xRangeMin; xMin < xRangeMax; xMin += q.config.Gen.DegreePerTile {
		for yMin := yRangeMin; yMin < yRangeMax; yMin += q.config.Gen.DegreePerTile {
			select {
			case <-done: // Check if we should stop due to an error
				return fmt.Errorf("export image stopped due to previous error")
			default:
				wg.Add(1)
				go func(x, y int) {
					defer wg.Done()
					err := q.ExportImageTile(projectPath,
						x, x+q.config.Gen.DegreePerTile,
						y, y+q.config.Gen.DegreePerTile,
						layerOutputName, layers)
					if err != nil {
						select {
						case errChan <- err:
						default:
						}
					}
				}(xMin, yMin)
			}
		}
	}

	wg.Wait()
	close(errChan)

	select {
	case <-done:
		return fmt.Errorf("export image failed")
	default:
		return nil
	}
}
