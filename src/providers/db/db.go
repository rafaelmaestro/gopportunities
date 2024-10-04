package db

import (
	"fmt"

	"github.com/fnunezzz/go-logger"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDatabase struct {
	DB *gorm.DB
}

func NewDatabase(config *config.Config) (*GormDatabase, error) {
	fmt.Println("config", config)
	sLog := logger.Get()

	dbConnectionRetries := config.Db.ConnectionRetries
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=America/Sao_Paulo", config.Db.Host, config.Db.Port, config.Db.User, config.Db.Name, config.Db.Pass)

	// TODO: Montar conexão à depender do driver (postgres, mysql, etc)
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	for dbConnectionRetries > 0 {
		if err != nil {
			sLog.Errorf("failed to connect to database, retrying %d, error: %s", dbConnectionRetries, err)
			dbConnectionRetries--
			db, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{})
		} else {
			break
		}

		if dbConnectionRetries == 0 {
			return nil, err
		}
	}

	// TODO: Migrate models (define where to put this)
	// db.AutoMigrate(&model.PrecoModel{})

	return &GormDatabase{
		DB: db,
	}, nil
}
