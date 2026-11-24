package main

import (
	"healthcheckservice/config"
	"healthcheckservice/logger"
	"healthcheckservice/monitor"
	"healthcheckservice/server"
)

func main() {
	// Initialise structured logger.
	logger.Init()

	// Load configuration from environment.
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Failed to load configuration: ", err)
		return
	}

	// Set up the health monitor.
	mon := monitor.New(cfg)
	mon.Start()
	defer mon.Stop()

	// Start the HTTP server.
	srv := server.New(cfg, mon)
	if err := srv.Start(); err != nil {
		logger.Error("Server terminated with error: ", err)
	}
}