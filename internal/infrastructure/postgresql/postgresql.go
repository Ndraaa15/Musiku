package postgresql

import (
	"log"

	"github.com/Ndraaa15/musiku/cmd/config"
	e "github.com/Ndraaa15/musiku/global/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func NewPostgreSqlClient() (*DB, error) {
	db, err := gorm.Open(postgres.Open(config.PostgresqlConfig()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("[musiku-postgresql] failed to connecting with musiku database : %v\n", err)
		return nil, e.ErrConnectDatabase
	}

	return &DB{db}, nil
}
