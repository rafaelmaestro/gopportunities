package preco

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/application/usecase"
	"github.com/rafaelmaestro/gopportunities/src/providers/akafka"
	"github.com/rafaelmaestro/gopportunities/src/providers/cache"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	httpServer "github.com/rafaelmaestro/gopportunities/src/providers/http"
)

type CriarPrecoRequestProps struct {
	Sku   int     `json:"sku"`
	Nome  string  `json:"nome"`
	Valor float64 `json:"valor"`
}

type CriarPrecoResponseProps struct {
	Id    string  `json:"id"`
	Sku   int     `json:"sku"`
	Nome  string  `json:"nome"`
	Valor float64 `json:"valor"`
}

type PrecoController struct {
	cfg               *config.Config
	httpServer        *httpServer.HttpServer
	kafkaProducer     akafka.IKafkaProducer
	criarPrecoUseCase usecase.ICriarPrecoUseCase
	cacheClient       cache.ICacheClient
}

func NewPrecoController(
	cfg *config.Config,
	httpServer *httpServer.HttpServer,
	kafkaProducer akafka.IKafkaProducer,
	usecase usecase.ICriarPrecoUseCase,
	cacheClient cache.ICacheClient,
) *PrecoController {

	httpPrecoGroup := httpServer.AppGroup.Group("/preco")

	httpServer.AppGroup = httpPrecoGroup

	controller := &PrecoController{
		cfg:               cfg,
		httpServer:        httpServer,
		kafkaProducer:     kafkaProducer,
		criarPrecoUseCase: usecase,
		cacheClient:       cacheClient,
	}

	controller.registerRoutes()
	controller.registerEventListeners()

	return controller
}

func (precoController PrecoController) registerRoutes() {
	precoController.httpServer.AppGroup.GET("/health", precoController.HealthCheck)
	precoController.httpServer.AppGroup.GET("/teste", precoController.Teste)
	precoController.httpServer.AppGroup.POST("", precoController.CriarPreco)
	precoController.httpServer.AppGroup.GET("/redis", precoController.TesteRedis)

}

func (precoController PrecoController) registerEventListeners() {
	// precoController.Teste2()
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
	go akafka.NewKafkaConsumer(precoController.cfg, &akafka.AKafkaConsumerConfig{
		ConsumerGroup:     precoController.cfg.Kafka.GroupID + "-preco",
		Topic:             precoController.cfg.Kafka.Topics["TESTE_TOPIC"],
		ConcurrentReaders: precoController.cfg.Kafka.ConcurrentReaders,
		MessageDto:        reflect.TypeOf(usecase.CriarPrecoDto{}), // Modificação aqui
		Handle: func(ctx context.Context, dto interface{}) error {
			precoController.criarPrecoUseCase.Execute(dto.(*usecase.CriarPrecoDto)) // Conversão de tipo
			return nil
		},
	})

	return nil
}

func (precoController PrecoController) TesteRedis(pctx echo.Context) error {
	err := precoController.cacheClient.Set(pctx.Request().Context(), "teste", "teste", 0)
	if err != nil {
		fmt.Println(err)
		return pctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	testeKey, err := precoController.cacheClient.Get(pctx.Request().Context(), "teste")

	if err != nil {
		fmt.Println(err)
		return pctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	fmt.Println(testeKey)
	return pctx.JSON(http.StatusOK, map[string]string{"status": testeKey})
}

func (precoController PrecoController) CriarPreco(pctx echo.Context) error {
	var userRequest CriarPrecoRequestProps

	if err := pctx.Bind(&userRequest); err != nil {
		return pctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	dto := &usecase.CriarPrecoDto{
		Sku:   userRequest.Sku,
		Nome:  userRequest.Nome,
		Valor: userRequest.Valor,
	}

	preco, err := precoController.criarPrecoUseCase.Execute(dto)

	if err != nil {
		return pctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response := CriarPrecoResponseProps{
		Id:    preco.GetId(),
		Sku:   preco.GetProduto().GetSKU(),
		Nome:  preco.GetProduto().GetNome(),
		Valor: preco.GetValor(),
	}

	return pctx.JSON(http.StatusOK, response)
}
