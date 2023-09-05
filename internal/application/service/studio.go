package service

import "github.com/Ndraaa15/musiku/internal/domain/repository"

type StudioService struct {
	StudioRepository repository.StudioRepositoryImpl
}

func NewStudioService(studioRepository repository.StudioRepositoryImpl) *StudioService {
	return &StudioService{
		StudioRepository: studioRepository,
	}
}
