package repository

import (
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
)

type StudioRepository struct {
	db *mysql.DB
}

func NewStudioRepository(db *mysql.DB) *StudioRepository {
	return &StudioRepository{
		db: db,
	}
}
