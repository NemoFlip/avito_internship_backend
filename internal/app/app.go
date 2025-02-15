package app

import (
	"avito_internship_backend/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	r := gin.Default()

	URL := fmt.Sprintf(":%s", cfg.Port)
	r.Run(URL)
}
