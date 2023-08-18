package postgresql

import (
	"errors"
	"log"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Instrument{},
		&entity.Studio{},
		&entity.Venue{},
	); err != nil {
		log.Fatalf("[musiku-postgresql] failed to migrate musiku database : %v\n", err)
		return errors.New("failed to migrate musiku database : " + err.Error())
	}

	return nil
}
