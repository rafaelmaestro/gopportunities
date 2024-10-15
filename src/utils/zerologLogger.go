package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func ZerologLogger() *zerolog.Logger {
	// Retrieve the current environment from the environment variable
	appEnv := os.Getenv("APP_ENV")

	// If the environment is "development", configure the logger with colors
	if appEnv == "development" {
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			FormatLevel: func(i interface{}) string {
				// Add color to log levels
				if ll, ok := i.(string); ok {
					switch ll {
					case "debug":
						return "\033[36mDEBUG\033[0m" // Cyan
					case "info":
						return "\033[32mINFO\033[0m" // Green
					case "warn":
						return "\033[33mWARN\033[0m" // Yellow
					case "error":
						return "\033[31mERROR\033[0m" // Red
					case "fatal":
						return "\033[35mFATAL\033[0m" // Magenta
					case "panic":
						return "\033[41mPANIC\033[0m" // Red background
					default:
						return ll
					}
				}
				return "UNKNOWN"
			},
			FormatMessage: func(i interface{}) string {
				// Add color to the actual log message
				return fmt.Sprintf("\033[34m%s\033[0m", i) // Blue
			},
			FormatFieldName: func(i interface{}) string {
				// Make field names bold
				return fmt.Sprintf("\033[1m%s:\033[0m", i)
			},
			FormatFieldValue: func(i interface{}) string {
				// Make field values yellow
				return fmt.Sprintf("\033[33m%s\033[0m", i)
			},
		}

		// Create and return a zerolog logger with the custom ConsoleWriter
		zl := zerolog.New(consoleWriter).With().Timestamp().Logger()
		return &zl
	}

	// For non-development environments, return the default zerolog logger
	zl := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &zl
}
