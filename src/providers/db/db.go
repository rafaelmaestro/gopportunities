package db

import (
	"fmt"
	"log/slog"

	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDatabase struct {
	DB *gorm.DB
}

func NewDatabase(config *config.Config) (*GormDatabase, error) {
	dbConnectionRetries := config.Db.ConnectionRetries
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=America/Sao_Paulo", config.Db.Host, config.Db.Port, config.Db.User, config.Db.Name, config.Db.Pass)

	// TODO: Montar conexão à depender do driver
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	for dbConnectionRetries > 0 {
		if err != nil {
			slog.Error("Failed to connect to database, retrying...", "error", err)
			dbConnectionRetries--
			db, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{})
		} else {
			break
		}
	}

	// TODO: Migrate models (define where to put this)
	// slog.Debug("Migrating database...")
	// db.AutoMigrate(&model.PrecoModel{})

	return &GormDatabase{
		DB: db,
	}, nil
}
