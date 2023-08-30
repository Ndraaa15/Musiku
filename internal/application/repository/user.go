package repository

import (
	"context"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/domain/repository"
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *mysql.DB
}

func NewUserRepository(db *mysql.DB) repository.UserRepositoryImpl {
	return &UserRepository{db}
}

func (ur *UserRepository) Create(user *entity.User, ctx context.Context) (*entity.User, error) {
	if err := ur.db.Debug().WithContext(ctx).Create(&user).Error; err != nil {
		return nil, gorm.ErrRegistered
	}
	return user, nil
}
