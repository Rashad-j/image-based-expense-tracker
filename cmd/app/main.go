package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Rashad-j/image-based-expense-tracker/internal/api/handlers"
	"github.com/Rashad-j/image-based-expense-tracker/internal/api/routes"
	"github.com/Rashad-j/image-based-expense-tracker/internal/core/ocr"
	"github.com/Rashad-j/image-based-expense-tracker/internal/core/parser"
	"github.com/Rashad-j/image-based-expense-tracker/internal/services"
	"github.com/Rashad-j/image-based-expense-tracker/pkg/chatgpt"
	"github.com/Rashad-j/image-based-expense-tracker/pkg/config"
	"github.com/Rashad-j/image-based-expense-tracker/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.ReadEnvConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Init logger
	_ = logger.InitLogger(cfg.LoggerLevel)

	// Initialize OCR processor
	OCRProcessor := ocr.NewGoogleVisionProcessor(cfg)

	// Initialize parser
	chatgptClient := chatgpt.NewClientUsingOpenapiLib(cfg)
	TextParser := parser.NewChatGPTParser(cfg, chatgptClient)

	// Initialize expense service
	ExpenseService := services.NewExpenseService(*cfg, OCRProcessor, TextParser)

	// Initialize API handlers
	handlers := handlers.NewExpenseHandler(ExpenseService)

	// Initialize router
	router := routes.SetupRouter("tmp-key", handlers)

	// Start server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		log.Printf("Starting server on port %s...", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting.")
}
