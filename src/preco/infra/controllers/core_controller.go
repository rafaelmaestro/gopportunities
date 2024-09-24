package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafaelmaestro/gopportunities/src/preco/application/usecase"
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

type CoreController struct {
	// ...instancias de usecases
	context *gin.Context
	criarPrecoUseCase usecase.ICriarPrecoUseCase
}

func NewCoreController(criarPrecoUseCase usecase.ICriarPrecoUseCase) *CoreController {
	return &CoreController{
		criarPrecoUseCase: criarPrecoUseCase,
	}
}

func (s CoreController) HealthCheck() {
	s.context.JSON(200, gin.H{
		"message": "Hello, World!",
		"current_date": time.Now().Format(time.RFC3339),
	})
}


func (s CoreController) CriarPreco() {
	var userRequest CriarPrecoRequestProps

	if  err := s.context.ShouldBindJSON(&userRequest); err != nil {
		s.context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	preco, err := s.criarPrecoUseCase.Execute(userRequest.Sku, userRequest.Nome, userRequest.Valor)

	if err != nil {
		s.context.JSON(400, gin.H{"error": err.Error()})

		return
	}

	response := CriarPrecoResponseProps{
		Id: preco.GetId(),
		Sku: preco.GetProduto().GetSKU(),
		Nome: preco.GetProduto().GetNome(),
		Valor: preco.GetValor(),
	}

	s.context.JSON(200, response)

	fmt.Println(userRequest)
}
