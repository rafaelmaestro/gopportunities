package utils

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/rs/zerolog"
)

// Middleware para capturar trace automaticamente em cada requisição HTTP
func EchoTracer(logger zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Início da requisição
			start := time.Now()
			req := c.Request()
			res := c.Response()

			logger.Trace().
				Str("method", req.Method).
				Str("path", req.URL.Path).
				Str("host", req.Host).
				Str("remote_addr", c.RealIP()).
				Msg("Iniciando requisição")

			// Processar a requisição
			err := next(c)

			// Verificar o status final após a execução da requisição
			status := res.Status
			if err != nil && status == http.StatusOK {
				// Se um erro aconteceu e o status é 200, modifique o status para o erro correto
				if he, ok := err.(*echo.HTTPError); ok {
					status = he.Code
				} else {
					status = http.StatusInternalServerError
				}
			}

			// Utilizando o nível correto de log baseado no status final
			switch {
			case status >= 500:
				logger.Error(). // Vermelho por padrão para erros (500+)
						Str("method", req.Method).
						Str("path", req.URL.Path).
						Int("status", status).
						Dur("latency", time.Since(start)).
						Msg("Erro no processamento da requisição")

			case status == 404:
				logger.Warn(). // Amarelo/Laranja por padrão para 404
						Str("method", req.Method).
						Str("path", req.URL.Path).
						Int("status", status).
						Dur("latency", time.Since(start)).
						Msg("Rota não encontrada")

			default:
				logger.Info(). // Verde/Azul por padrão para status 200 e outras respostas de sucesso
						Str("method", req.Method).
						Str("path", req.URL.Path).
						Int("status", status).
						Dur("latency", time.Since(start)).
						Msg("Requisição processada com sucesso")
			}

			return err
		}
	}
}
