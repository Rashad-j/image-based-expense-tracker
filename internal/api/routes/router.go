package routes

import (
	"github.com/Rashad-j/image-based-expense-tracker/internal/api/handlers"
	"github.com/Rashad-j/image-based-expense-tracker/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes all application routes.
func SetupRouter(authKey string, handler *handlers.ExpenseHandler) *gin.Engine {
	router := gin.New()

	// Logger and Recovery middleware are enabled by default
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middleware.AuthMiddleware(authKey))

	router.POST("/v1/analyze", handler.AnalyzeReceipt)

	return router
}
