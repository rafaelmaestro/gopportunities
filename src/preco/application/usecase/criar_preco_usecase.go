package usecase

import (
	"fmt"

	"github.com/rafaelmaestro/gopportunities/src/preco/domain"
	"github.com/rafaelmaestro/gopportunities/src/preco/infra/repositories"
)

type CriarPrecoUseCase struct {
	// instancias de repositorios
	// instancias de servicos
	precoRepository repositories.IPrecoRepository
}

type ICriarPrecoUseCase interface {
	Execute(sku int, nome string, valor float64) (*domain.PrecoAggregate, error)
}

func NewCriarPrecoUseCase(precoRepository repositories.IPrecoRepository) *CriarPrecoUseCase {
	// construir instancia de usecase passando as instancias de repositorios e servicos
	return &CriarPrecoUseCase{
		precoRepository: precoRepository,
	}
}

func (u *CriarPrecoUseCase) Execute(sku int, nome string, valor float64) (*domain.PrecoAggregate, error) {
	produtoObject, err := domain.NewProduto(sku, nome)

	if err != nil {
		return nil,err
	}

	preco, err := domain.NewPreco(produtoObject, valor)

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
