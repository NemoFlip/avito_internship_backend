package main

import (
	"avito_internship_backend/config"
	"fmt"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg.Port)
	// TODO: initialize config and logger - then continue in internal/app/app.go
}
