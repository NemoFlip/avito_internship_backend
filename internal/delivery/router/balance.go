package router

import (
	"avito_internship_backend/internal/delivery/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterBalanceRoutes(r *gin.Engine) {
	balanceGroup := r.Group("/balance") // TODO: make check for authorization
	{
		balanceGroup.POST("/deposit", handlers.Deposit)
		balanceGroup.GET("", handlers.GetBalance)
	}
}
