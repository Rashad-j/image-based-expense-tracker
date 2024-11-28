package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExpenseHandler provides endpoints for handling expenses.
type ExpenseHandler struct {
	// config and database, etc
}

// NewExpenseHandler creates a new ExpenseHandler.
func NewExpenseHandler() *ExpenseHandler {
	return &ExpenseHandler{}
}

// AnalyzeReceipt handles receipt analysis requests.
func (h *ExpenseHandler) AnalyzeReceipt(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Receipt analyzed successfully"})
}
