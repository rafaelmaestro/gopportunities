package repositories

import (
	"errors"

	"github.com/rafaelmaestro/gopportunities/src/db"
	"github.com/rafaelmaestro/gopportunities/src/preco/domain"
	"github.com/rafaelmaestro/gopportunities/src/preco/infra/mappers"
)

var ErrFalhaAoPersistirObjeto = errors.New("erro ao persistir objeto no banco de dados")

type IPrecoRepository interface {
	Create(*domain.PrecoAggregate) error
}

type PrecoRepository struct {
	database *db.GormDatabase
}

func NewPrecoRepository(database *db.GormDatabase) *PrecoRepository {
	// construir instancia de repositorio passando as instancias de conexão com banco de dados
	return &PrecoRepository{
		database: &db.GormDatabase{},
	}
}

func (r *PrecoRepository) Create(preco *domain.PrecoAggregate) error {
	// converter de domain.PrecoAggregate para model.PrecoModel
	precoModel := mappers.ToModel(preco)

	if precoModel == nil {
		return ErrFalhaAoPersistirObjeto
	}

	result := r.database.DB.Create(precoModel)

	if result.Error != nil {
		return ErrFalhaAoPersistirObjeto
	}
	return nil
}
