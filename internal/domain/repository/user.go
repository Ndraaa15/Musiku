package repository

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
)

type UserRepositoryImpl interface {
	Create(user *entity.User, ctx context.Context) (*entity.User, error)
}
