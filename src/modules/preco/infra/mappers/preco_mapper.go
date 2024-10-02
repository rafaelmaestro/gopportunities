package mappers

import (
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/domain"
	model "github.com/rafaelmaestro/gopportunities/src/modules/preco/infra/models"
)

type IPrecoMapper interface {
	ToModel(*domain.PrecoAggregate) *model.PrecoModel
	ToDomain(*model.PrecoModel) *domain.PrecoAggregate
}

type PrecoMapper struct {
}

func NewPrecoMapper() *PrecoMapper {
	return &PrecoMapper{}
}

func (precoMapper *PrecoMapper) ToModel(domain *domain.PrecoAggregate) *model.PrecoModel {
	return &model.PrecoModel{
		ID:    domain.GetId(),
		Sku:   domain.GetProduto().GetSKU(),
		Nome:  domain.GetProduto().GetNome(),
		Valor: domain.GetValor(),
	}
}

func (precoMapper *PrecoMapper) ToDomain(model *model.PrecoModel) *domain.PrecoAggregate {
	produtoObject, err := domain.NewProduto(model.Sku, model.Nome)

	if err != nil {
		return nil
	}

	preco, err := domain.NewPreco(produtoObject, model.Valor)

	if err != nil {
		return nil
	}

	return preco
}
