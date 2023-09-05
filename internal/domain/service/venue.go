package service

import "github.com/Ndraaa15/musiku/internal/domain/entity"

type VenueServiceImpl interface {
	GetAllVenue() ([]*entity.Venue, error)
}
