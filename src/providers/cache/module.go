package cache

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module(
		"cache",
		fx.Provide(fx.Annotate(
			NewCacheClient, fx.As(new(ICacheClient)),
		)),
	)
}
