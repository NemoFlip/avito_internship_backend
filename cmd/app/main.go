package main

import (
	"avito_internship_backend/config"
	_ "avito_internship_backend/docs"
	"avito_internship_backend/internal/app"
	"avito_internship_backend/pkg/logging"
)

func main() {
	logger := logging.InitLogger()
	cfg, err := config.NewConfig()
	if err != nil {
		logger.ErrorLogger.Fatal(err)
	}
	logger.InfoLogger.Info("Config was initialized")

	app.Run(logger, cfg)
}
