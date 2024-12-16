package handlers

import (
	"net/http"

	"github.com/Rashad-j/image-based-expense-tracker/internal/services"
	"github.com/gin-gonic/gin"
)

// ExpenseHandler provides endpoints for handling expenses.
type ExpenseHandler struct {
	expenseService services.ExpenseService
}

// NewExpenseHandler creates a new ExpenseHandler.
func NewExpenseHandler(expenseService services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: expenseService,
	}
}

// AnalyzeReceipt handles receipt analysis requests.
func (h *ExpenseHandler) AnalyzeReceipt(c *gin.Context) {
	// read image bytes from request and get io.Reader to send to analyze
	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read image"})
		return
	}
	file, err := image.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to open image"})
		return
	}
	defer file.Close()

	expenses, err := h.expenseService.AnalyzeReceipt(c, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to analyze receipt"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"expenses": expenses})
}
