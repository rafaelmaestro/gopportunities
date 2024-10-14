package preco

import (
	"context"
	"log"

	"github.com/rafaelmaestro/gopportunities/src/modules/preco/application/usecase"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/infra/mappers"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/infra/repositories"
	"github.com/rafaelmaestro/gopportunities/src/providers/akafka"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"preco",
		/* Should initialize all the providers below (repositories, usecases, controllers, etc)
		Elements from outside the infrastructure layer shouldnt depends directly on implementations from the infrastructure layer
		Use interfaces to make the dependencies between layers more flexible (fx.As, fx.Annotate)
		*/
		fx.Provide(fx.Annotate(
			repositories.NewPrecoRepository, fx.As(new(repositories.IPrecoRepository)),
		)),
		fx.Provide(fx.Annotate(
			usecase.NewCriarPrecoUseCase, fx.As(new(usecase.ICriarPrecoUseCase)),
		)),
		fx.Provide(mappers.NewPrecoMapper),
		// fx.Provide(cache.NewCacheClient),

		// Should initialize the kafka producer and add a hook to close it on application shutdown
		// Close kafka producer on application shutdown, to avoid memory leaks
		fx.Provide(fx.Annotate(akafka.NewKafkaProducer, fx.As(new(akafka.IKafkaProducer)))),
		fx.Invoke(func(lifecycle fx.Lifecycle, producer akafka.IKafkaProducer) {
			lifecycle.Append(fx.Hook{
				OnStop: func(ctx context.Context) error { // Recebe o contexto aqui
					if err := producer.Close(); err != nil {
						log.Fatalf("Failed to close producer: %v", err)
					}
					return nil
				},
			})
		}),

		// Should initialize the controllers and call the registerRoutes and registerEventListeners methods with fx.Invoke
		// fx.Provide(controllers.NewPrecoController),
		// fx.Invoke(func(precoController *controllers.PrecoController) {
		//     precoController.RegisterRoutes()
		// 	precoController.RegisterEventListeners()
		// }),

		// Cleaner way to do the above commented code
		fx.Invoke(NewPrecoController),
	)
}
