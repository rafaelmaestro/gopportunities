package db

import (
	"fmt"
	"time"

	"github.com/fnunezzz/go-logger"
	"github.com/rafaelmaestro/gopportunities/src/providers/config"
	"github.com/rafaelmaestro/gopportunities/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDatabase struct {
	DB *gorm.DB
}

func NewDatabase(config *config.Config) (*GormDatabase, error) {
	sLog := logger.Get()

	// Definir valor padrão para ConnectionRetries se não estiver definido
	dbConnectionRetries := config.Db.ConnectionRetries
	if dbConnectionRetries == 0 {
		dbConnectionRetries = 3 // Valor padrão, por exemplo, 3 tentativas
	}

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=America/Sao_Paulo", config.Db.Host, config.Db.Port, config.Db.User, config.Db.Name, config.Db.Pass)
	var db *gorm.DB
	var err error

	// TODO: Montar conexão à depender do driver (postgres, mysql, etc)
	for i := 0; i < dbConnectionRetries; i++ {
		db, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{
			Logger: utils.LoggerGorm{Log: *utils.ZerologLogger()},
		})

		if err == nil {
			break
		}

		time.Sleep(5 * time.Second)
	}

	if err != nil {
		sLog.Errorf("failed to connect to database, error: %s", err)
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		sLog.Errorf("failed to get sql.DB from gorm.DB, error: %s", err)
		return nil, err
	}

	sqlDb.SetMaxIdleConns(config.Db.MaxIdleConnections)
	sqlDb.SetMaxOpenConns(config.Db.MaxOpenConnections)
	sqlDb.SetConnMaxLifetime(config.Db.ConnectionMaxLifetime)

	// TODO: Migrate models (define where to put this)
	// db.AutoMigrate(&model.PrecoModel{})

	// Adding a defer to close the connection when the function ends
	// TODO: Check if this is necessary
	// defer func() {
	// 	sqlDb, err := db.DB()
	// 	if err == nil {
	// 		sqlDb.Close()
	// 	}
	// }()

	return &GormDatabase{
		DB: db,
	}, nil
}
