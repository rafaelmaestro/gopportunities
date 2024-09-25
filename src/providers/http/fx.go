package httpServer

import (
	"fmt"

	"go.uber.org/fx"
)

func Module() fx.Option {
	fmt.Println("1 - httpServer.Module")
	return fx.Module(
		"http",
		fx.Provide(NewServer),
	)
}
