package mysql

import (
	"log"

	e "github.com/Ndraaa15/musiku/global/errors"
	"github.com/Ndraaa15/musiku/internal/domain/entity"
)

func Migration(db *DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Instrument{},
		&entity.Studio{},
		&entity.Venue{},
	); err != nil {
		log.Fatalf("[musiku-postgresql] failed to migrate musiku database : %v\n", err)
		return e.ErrMigrateDatabase
	}

	return nil
}
