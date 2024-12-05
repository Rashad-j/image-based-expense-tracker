package routes

import (
	"github.com/Rashad-j/image-based-expense-tracker/internal/api/handlers"
	"github.com/Rashad-j/image-based-expense-tracker/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes all application routes.
func SetupRouter(authKey string) *gin.Engine {
	router := gin.Default()

	// Middleware
	// router.Use(middleware.MetricsMiddleware())
	router.Use(middleware.AuthMiddleware(authKey))

	// Handlers
	expenseHandler := handlers.NewExpenseHandler()

	// Routes
	router.POST("/v1/analyze", expenseHandler.AnalyzeReceipt)

	return router
}
