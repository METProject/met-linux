package qgis

import (
	"fmt"
	"path/filepath"
)

// calculateTile calculates the tile name based on coordinates
func calculateTile(x, y int) string {
	var ns, ew string
	if y >= 0 {
		ns = "n"
	} else {
		ns = "s"
		y = -y
	}
	if x >= 0 {
		ew = "e"
	} else {
		ew = "w"
		x = -x
	}
	return fmt.Sprintf("%s%d%s%d", ns, y, ew, x)
}

// getTileOutputPath returns the full path for a tile image
func getTileOutputPath(outputFolder, tile, layerOutputName string) string {
	tileFolder := filepath.Join(outputFolder, fmt.Sprintf("%s", tile))
	if layerOutputName == "" {
		return filepath.Join(tileFolder, fmt.Sprintf("%s.png", tile))
	}
	return filepath.Join(tileFolder, fmt.Sprintf("%s_%s.png", tile, layerOutputName))
}
