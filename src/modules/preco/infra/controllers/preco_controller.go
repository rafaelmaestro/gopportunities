package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/application/usecase"
	"github.com/rafaelmaestro/gopportunities/src/providers/akafka"
	httpServer "github.com/rafaelmaestro/gopportunities/src/providers/http"
)

type CriarPrecoRequestProps struct {
	Sku   int     `json:"sku"`
	Nome  string  `json:"nome"`
	Valor float64 `json:"valor"`
}

type CriarPrecoResponseProps struct {
	Id string `json:"id"`
	Sku int `json:"sku"`
	Nome string `json:"nome"`
	Valor float64 `json:"valor"`
}

type PrecoController struct {
	httpServer *httpServer.HttpServer
	kafkaProducer akafka.IKafkaProducer
	criarPrecoUseCase usecase.ICriarPrecoUseCase
}

func NewPrecoController(
	httpServer *httpServer.HttpServer,
	kafkaProducer akafka.IKafkaProducer,
	usecase usecase.ICriarPrecoUseCase,
) *PrecoController {

	httpPrecoGroup := httpServer.AppGroup.Group("/preco")

	httpServer.AppGroup = httpPrecoGroup

	return &PrecoController{
		httpServer: httpServer,
		kafkaProducer: kafkaProducer,
		criarPrecoUseCase: usecase,
	}
}

func (precoController PrecoController) RegisterRoutes() {
	precoController.httpServer.AppGroup.GET("/health", precoController.HealthCheck)
	precoController.httpServer.AppGroup.GET("/teste", precoController.Teste)
	precoController.httpServer.AppGroup.GET("/teste2", precoController.Teste2)
}

func (precoController PrecoController) HealthCheck(pctx echo.Context) error {
	precoController.httpServer.AppGroup.GET("/health", func(pctx echo.Context) (err error) {
		return pctx.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})
	return pctx.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}

func (precoController PrecoController) Teste(pctx echo.Context) error {
	precoController.kafkaProducer.SendMessage("test", "teste", "teste")
	return pctx.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}

func (precoController PrecoController) Teste2(
	pctx echo.Context,
) error {
	precoController.kafkaProducer.SendMessage("test2", "teste2", "teste2")
	return pctx.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}


// func (s CoreController) CriarPreco() {
// 	var userRequest CriarPrecoRequestProps

	// if  err := s.context.ShouldBindJSON(&userRequest); err != nil {
	// 	s.context.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	// preco, err := s.criarPrecoUseCase.Execute(userRequest.Sku, userRequest.Nome, userRequest.Valor)

	// if err != nil {
	// 	s.context.JSON(400, gin.H{"error": err.Error()})

	// 	return
	// }

	// response := CriarPrecoResponseProps{
	// 	Id: preco.GetId(),
	// 	Sku: preco.GetProduto().GetSKU(),
	// 	Nome: preco.GetProduto().GetNome(),
	// 	Valor: preco.GetValor(),
	// }

	// s.context.JSON(200, response)

// 	fmt.Println(userRequest)
// }
