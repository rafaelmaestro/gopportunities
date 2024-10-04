package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/application/usecase"
	"github.com/rafaelmaestro/gopportunities/src/providers/akafka"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
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
	cfg *config.Config
	httpServer *httpServer.HttpServer
	kafkaProducer akafka.IKafkaProducer
	criarPrecoUseCase usecase.ICriarPrecoUseCase
}

func NewPrecoController(
	cfg *config.Config,
	httpServer *httpServer.HttpServer,
	kafkaProducer akafka.IKafkaProducer,
	usecase usecase.ICriarPrecoUseCase,
) *PrecoController {
	httpPrecoGroup := httpServer.AppGroup.Group("/preco")

	httpServer.AppGroup = httpPrecoGroup

	return &PrecoController{
		cfg: cfg,
		httpServer: httpServer,
		kafkaProducer: kafkaProducer,
		criarPrecoUseCase: usecase,
	}
}

func (precoController PrecoController) RegisterRoutes() {
	precoController.httpServer.AppGroup.GET("/health", precoController.HealthCheck)
	precoController.httpServer.AppGroup.GET("/teste", precoController.Teste)
}

func (precoController PrecoController) RegisterEventListeners() {
	precoController.Teste2()
}

func (precoController PrecoController) HealthCheck(pctx echo.Context) error {
	precoController.httpServer.AppGroup.GET("/health", func(pctx echo.Context) (err error) {
		return pctx.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})
	return pctx.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}

func (precoController PrecoController) Teste(pctx echo.Context) error {
	precoController.kafkaProducer.SendMessage(pctx.Request().Context(), "test", "teste", "teste", nil)
	return pctx.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}

func (precoController PrecoController) Teste2() error {
	kafkaConsumerConfig := &akafka.AKafkaConsumerConfig{
		ConsumerGroup: precoController.cfg.Kafka.GroupID + "-preco",
		Topic: precoController.cfg.Kafka.Topics["Teste2"],
		ConcurrentReaders: precoController.cfg.Kafka.ConcurrentReaders,
		Handle: func(ctx context.Context, msg *akafka.AKafkaMessage) error {
			fmt.Printf("[Handler Function] Received message at offset %d: %s = %s\n", msg.AMessage.Offset, string(msg.AMessage.Key), string(msg.AMessage.Value))
			return nil
		},
	}
	go akafka.NewKafkaConsumer(precoController.cfg, kafkaConsumerConfig)


	return nil
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
