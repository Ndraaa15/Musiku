package repository

import (
	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
)

type VenueRepository struct {
	db *mysql.DB
}

func NewVenueRepository(db *mysql.DB) repository.VenueRepositoryImpl {
	return &VenueRepository{db}
}

func (v *VenueRepository) GetAll() ([]*entity.Venue, error) {
	return nil, nil
}

func (v *VenueRepository) UpdateDay() (*entity.Venue, error) {
	return nil, nil
}
