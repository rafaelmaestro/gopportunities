package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func ZerologLogger() *zerolog.Logger {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "development" {
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			FormatLevel: func(i interface{}) string {
				if ll, ok := i.(string); ok {
					switch ll {
					case "debug":
						return "\033[36mDEBUG\033[0m" // Ciano
					case "info":
						return "\033[32mINFO\033[0m" // Verde
					case "warn":
						return "\033[33mWARN\033[0m" // Amarelo
					case "error":
						return "\033[31mERROR\033[0m" // Vermelho
					case "fatal":
						return "\033[35mFATAL\033[0m" // Magenta
					case "panic":
						return "\033[41mPANIC\033[0m" // Fundo vermelho
					default:
						return ll
					}
				}
				return "UNKNOWN"
			},
			FormatMessage: func(i interface{}) string {
				// Adiciona cor à mensagem de log
				return fmt.Sprintf("\033[34m%s\033[0m", i) // Azul
			},
			FormatFieldName: func(i interface{}) string {
				// Nomes de campos em negrito
				return fmt.Sprintf("\033[1m%s:\033[0m", i)
			},
			FormatFieldValue: func(i interface{}) string {
				// Valores dos campos em amarelo
				return fmt.Sprintf("\033[33m%s\033[0m", i)
			},
		}

		zl := zerolog.New(consoleWriter).With().Timestamp().Logger()
		return &zl
	}

	zl := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &zl
}
