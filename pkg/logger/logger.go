package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// InitLogger initializes the zerolog logger with the configured log level.
func InitLogger(logLevel string) zerolog.Logger {
	level, err := zerolog.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		fmt.Printf("Invalid log level '%s', defaulting to 'info'\n", logLevel)
		level = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(level)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Info().Str("log_level", level.String()).Msg("Logger initialized")
	return logger
}
