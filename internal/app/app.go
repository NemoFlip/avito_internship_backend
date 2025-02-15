package app

import (
	"avito_internship_backend/config"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Avito Internship Backend
// @description This is backend for the avito's microservise task
// @host localhost:8080
// @BasePath /
func Run(cfg *config.Config) {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	url := fmt.Sprintf(":%s", cfg.Port)
	if err := r.Run(url); err != nil {
		panic(err)
	}
}
