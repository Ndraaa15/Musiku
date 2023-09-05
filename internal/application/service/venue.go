package service

import (
	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
	"github.com/Ndraaa15/musiku/internal/domain/service"
	"github.com/gofrs/uuid"
)

type VenueService struct {
	VenueRepository repository.VenueRepositoryImpl
}

func NewVenueService(VenueRepository repository.VenueRepositoryImpl) service.VenueServiceImpl {
	return &VenueService{VenueRepository}
}

func (v *VenueService) GetAllVenue() ([]*entity.Venue, error) {
	return nil, nil
}

func (v *VenueService) RentVenue(id uuid.UUID) (*entity.Venue, error) {
	return nil, nil
}
