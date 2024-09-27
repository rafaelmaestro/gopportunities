package preco

import (
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/application/usecase"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/infra/controllers"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/infra/repositories"
	"github.com/rafaelmaestro/gopportunities/src/providers/akafka"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"go.uber.org/fx"
)

// func createCoreRouter(server http.HttpServer, controller *controllers.CoreController) {
//     server.RegisterGroup("/preco", func(group http.RouteGroup) {
//         group.RegisterRoute("POST", "", controller.CriarPreco) // O contexto é passado como interface
//         // Adicione mais rotas aqui, se necessário
//     })
// }

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
		fx.Provide(fx.Annotate(akafka.NewKafkaProducer, fx.As(new(akafka.IKafkaProducer)))),

		// Should initialize the controllers and call the registerRoutes methods with fx.Invoke
        fx.Provide(controllers.NewPrecoController),
        fx.Invoke(func(precoController *controllers.PrecoController) {
            precoController.RegisterRoutes()
        }),

		// Should configure the AKafkaConsumer struct providing the consumerGroup and topics
		// This turns the AKafkaConsumer struct into a dependency that can be injected into the NewKafkaConsumer function
		// Doing this, we can create consumers with different configurations for different modules
		fx.Provide(func() *akafka.AKafkaConsumer {
			return &akafka.AKafkaConsumer{
				ConsumerGroup: "gopportunities-preco",   // Defina o nome do grupo de consumidores
				Topics: []string{"test", "test2"},  // Defina os tópicos que deseja consumir
			}
		}),
		fx.Invoke(func(config *config.Config, consumer *akafka.AKafkaConsumer) {
			akafka.NewKafkaConsumer(config, consumer)
		}),
	)
}
