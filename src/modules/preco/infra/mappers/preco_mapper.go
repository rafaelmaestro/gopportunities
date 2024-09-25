package mappers

import (
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/domain"
	model "github.com/rafaelmaestro/gopportunities/src/modules/preco/infra/models"
)

func ToModel(domain *domain.PrecoAggregate) *model.PrecoModel {
	return &model.PrecoModel{
		ID:    domain.GetId(),
		Sku:   domain.GetProduto().GetSKU(),
		Nome:  domain.GetProduto().GetNome(),
		Valor: domain.GetValor(),
	}
}

func ToDomain(model *model.PrecoModel) *domain.PrecoAggregate {
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
