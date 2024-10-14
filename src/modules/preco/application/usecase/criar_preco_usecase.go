package usecase

import (
	"fmt"

	"github.com/rafaelmaestro/gopportunities/src/modules/preco/domain"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/infra/repositories"
)


type CriarPrecoDto struct {
	Sku   int     `json:"sku"`
	Nome  string  `json:"nome"`
	Valor float64 `json:"valor"`
}

type CriarPrecoUseCase struct {
	// instancias de repositorios
	// instancias de servicos
	precoRepository repositories.IPrecoRepository
}

type ICriarPrecoUseCase interface {
	Execute(props *CriarPrecoDto) (*domain.PrecoAggregate, error)
}

func NewCriarPrecoUseCase(precoRepository repositories.IPrecoRepository) *CriarPrecoUseCase {
	// construir instancia de usecase passando as instancias de repositorios e servicos
	return &CriarPrecoUseCase{
		precoRepository: precoRepository,
	}
}

func (u *CriarPrecoUseCase) Execute(props *CriarPrecoDto) (*domain.PrecoAggregate, error) { // TODO: return dto


	fmt.Println("CriarPrecoUseCase.Execute")
	fmt.Print("%v", props)

	produtoObject, err := domain.NewProduto(props.Sku, props.Nome)

	if err != nil {
		return nil,err
	}

	preco, err := domain.NewPreco(produtoObject, props.Valor)

	if err != nil {
		return nil,err
	}

	err = u.precoRepository.Create(preco)

	if err != nil {
		return nil,err
	}

	// TODO: testar implementação com kafka enviando mensagem de preco criado

	fmt.Println(preco)

	return preco, nil
}
