package preco

import (
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/application/usecase"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/infra/controllers"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/infra/repositories"
	"github.com/rafaelmaestro/gopportunities/src/providers/akafka"
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
		// fx.Invoke(akafka.NewKafkaProducer),

		// fx.Invoke(akafka.NewKafkaConsumer),
		fx.Invoke(controllers.HealthCheck),
		fx.Invoke(controllers.Teste),

		// Should initialize all the router groups using Invoke	below
		// fx.Invoke(createPrecoRouter),
	)
}
