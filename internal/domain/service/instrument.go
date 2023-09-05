package service

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
)

type InstrumentServiceImpl interface {
	GetAllInstrument(ctx context.Context) ([]*entity.Instrument, error)
	GetByID(ctx context.Context, id uint) (*entity.Instrument, error)
	RentInstrument(ctx context.Context, id uint) (*entity.Instrument, error)
}
