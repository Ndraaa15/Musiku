package postgresql

import (
	"errors"
	"log"

	"github.com/Ndraaa15/musiku/cmd/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func NewPostgreClient() (*DB, error) {
	db, err := gorm.Open(postgres.Open(config.PostgresqlConfig()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("[musiku-postgresql] failed to connecting with musiku database : %v\n", err)
		return nil, errors.New("failed to connecting with musiku database : " + err.Error())
	}

	return &DB{db}, nil
}
