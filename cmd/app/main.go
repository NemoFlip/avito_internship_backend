package main

import (
	"avito_internship_backend/config"
	_ "avito_internship_backend/docs"
	"avito_internship_backend/internal/app"
	"avito_internship_backend/pkg/logger"
	"fmt"
)

func main() {
	logging := logger.InitLogger()
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	logging.InfoLogger.Info("Cfg initialized")

	app.Run(cfg)
	// TODO: initialize ogger - then continue in internal/app/app.go
}
