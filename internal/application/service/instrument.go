package service

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
)

type InstrumentService struct {
	InstrumentRepository repository.InstrumentRepositoryImpl
}

func NewInstrumentService(instrumentRepository repository.InstrumentRepositoryImpl) *InstrumentService {
	return &InstrumentService{
		InstrumentRepository: instrumentRepository,
	}
}

func (is *InstrumentService) GetAllInstrument(ctx context.Context) ([]*entity.Instrument, error) {
	instruments, err := is.InstrumentRepository.GetAllInstrument(ctx)
	if err != nil {
		return nil, err
	}
	return instruments, nil
}

func (is *InstrumentService) GetByID(ctx context.Context, id uint) (*entity.Instrument, error) {
	instrument, err := is.InstrumentRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return instrument, nil
}

func (is *InstrumentService) RentInstrument(ctx context.Context, id uint) (*entity.Instrument, error) {
	instrument, err := is.InstrumentRepository.Update(ctx, &entity.Instrument{}, id)
	if err != nil {
		return nil, err
	}
	return instrument, nil
}
