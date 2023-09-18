package mysql

import (
	"log"

	"github.com/Ndraaa15/musiku/global/errors"
	"github.com/Ndraaa15/musiku/internal/domain/entity"
)

func Migration(db *DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Instrument{},
		&entity.RentInstrument{},
		&entity.Studio{},
		&entity.RentStudio{},
		&entity.Time{},
		&entity.StartTime{},
		&entity.EndTime{},
		&entity.Venue{},
		&entity.VenueDay{},
		&entity.ApplyVenue{},
		&entity.Day{},
	); err != nil {
		log.Fatalf("[musiku-postgresql] failed to migrate musiku database : %v\n", err)
		return errors.ErrMigrateDatabase
	}
	return nil
}
