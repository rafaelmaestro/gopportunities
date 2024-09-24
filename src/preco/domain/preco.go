package domain

import (
	"errors"

	"github.com/google/uuid"
)

var ErrPrecoValorInvalido = errors.New("valor do produto deve ser maior que zero")
var ErrPrecoValorMuitoAlto = errors.New("valor do produto deve ser menor que 10000")

type PrecoAggregate struct {
	id uuid.UUID
	produto ProdutoEntity
	valor float64
	// other props to calculate the best price for the product
}

func NewPreco(produto *ProdutoEntity, valor float64) (*PrecoAggregate, error) {
	if valor <= 0 {
		return nil, ErrPrecoValorInvalido
	}

	if valor > 10000 {
		return nil, ErrPrecoValorMuitoAlto
	}

	id := uuid.New()

	if produto == nil {
		return nil, errors.New("produto não pode ser nulo")
	}

	preco := &PrecoAggregate{
		id: id,
		valor: valor,
		produto: *produto,
	}

	return preco, nil
}

func (p *PrecoAggregate) GetId() string {
	return p.id.String()
}

func (p *PrecoAggregate) GetProduto() ProdutoEntity {
	return p.produto
}

func (p *PrecoAggregate) GetValor() float64 {
	return p.valor
}