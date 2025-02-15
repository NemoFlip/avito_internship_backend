package main

import (
	"avito_internship_backend/config"
	"avito_internship_backend/internal/app"
	"fmt"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	app.Run(cfg)
	// TODO: initialize ogger - then continue in internal/app/app.go
}
