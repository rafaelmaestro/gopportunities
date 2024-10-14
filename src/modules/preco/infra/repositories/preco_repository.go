package repositories

import (
	"errors"

	"github.com/rafaelmaestro/gopportunities/src/modules/preco/domain"
	"github.com/rafaelmaestro/gopportunities/src/modules/preco/infra/mappers"
	"github.com/rafaelmaestro/gopportunities/src/providers/db"
)

var ErrFalhaAoPersistirObjeto = errors.New("erro ao persistir objeto no banco de dados")

type IPrecoRepository interface {
	Create(*domain.PrecoAggregate) error
}

type PrecoRepository struct {
	database *db.GormDatabase
	mapper   *mappers.PrecoMapper
}

func NewPrecoRepository(database *db.GormDatabase, mapper *mappers.PrecoMapper) *PrecoRepository {
	return &PrecoRepository{
		database: database,
		mapper:   mapper,
	}
}

func (r *PrecoRepository) Create(preco *domain.PrecoAggregate) error {
	precoModel := r.mapper.ToModel(preco)

	if precoModel == nil {
		return ErrFalhaAoPersistirObjeto
	}

	result := r.database.DB.Create(precoModel)

	if result.Error != nil {
		return ErrFalhaAoPersistirObjeto
	}
	return nil
}
