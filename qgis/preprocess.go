package qgis

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/imide/met-linux/util"
	"go.uber.org/zap"
)

func formatCmd(cmd, pbfPath, outputPath string) string {
	return strings.Replace(strings.Replace(cmd, "{P_FILE}", pbfPath, -1), "{OUT}", outputPath, -1)
}

func cmdRunner(name, cmd string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()

	start := time.Now()
	zap.L().Info(fmt.Sprintf("Running %s...", name))
	execCmd := exec.Command("bash", "-c", cmd)
	err := execCmd.Run()
	duration := time.Since(start)

	if err != nil {
		zap.L().Error(fmt.Sprintf("Failed to run %s: %s", name, err))
		<-sem
		return
	}

	zap.L().Info(fmt.Sprintf("Finished running %s in %s", name, duration))
	<-sem
}

func preprocessOSM(cfg *util.Config) {
	zap.L().Info("Preprocessing OSM...")

	if _, err := os.Stat(cfg.Path.OutputFolder); os.IsNotExist(err) {
		os.MkdirAll(cfg.Path.OutputFolder, os.ModePerm)
	}

	// build a buffered semaphore channel
	sem := make(chan struct{}, cfg.Gen.Threads)
	var wg sync.WaitGroup

	// run the commands

}
