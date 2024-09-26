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

func HealthCheck(
	server *httpServer.HttpServer,
) any {
	precoGroup := server.Group("/preco")

	precoGroup.GET("/health", func(pctx echo.Context) (err error) {
		return pctx.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})
	return nil
}

func Teste(
	server *httpServer.HttpServer,
	kafkaProducer akafka.IKafkaProducer,
	usecase usecase.ICriarPrecoUseCase,
) any {
	precoGroup := server.Group("/preco")

	precoGroup.GET("/teste", func(pctx echo.Context) (err error) {
		kafkaProducer.SendMessage("test", "teste", "teste")
		return pctx.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})
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
