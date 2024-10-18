package main

import (
	"github.com/imide/met-linux/util"
	"go.uber.org/zap"
)

func main() {
	// Setup logger
	util.NewLogger()
	zap.L().Info("MET Linux is starting...")

	// TODO: Verify config
	zap.L().Info("Loading config...")
	cfg, err := util.LoadConfig("config.toml")
	if err != nil {
		zap.L().Fatal("Failed to load config", zap.Error(err))
	}
	zap.L().Info("Config loaded successfully")

	// TODO: Run the generator
	zap.L().Info("Generator starting...")

}
