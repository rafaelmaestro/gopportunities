package domain

import "errors"

var ErrProdutoSkuInvalido = errors.New("sku deve estar preenchido e ser maior que zero")
var ErrProdutoNomeInvalido = errors.New("nome do produto deve estar preenchido")

type ProdutoEntity struct {
	sku  int
	nome string
}

func NewProduto(sku int, nome string) (*ProdutoEntity, error) {

	if sku <= 0 {
		return nil, ErrProdutoSkuInvalido
	}

	if nome == "" {
		return nil, ErrProdutoNomeInvalido
	}

	instance := &ProdutoEntity{
		sku:  sku,
		nome: nome,
	}

	return instance, nil

}

func (p ProdutoEntity) GetSKU() int {
	return p.sku
}

func (p ProdutoEntity) GetNome() string {
	return p.nome
}