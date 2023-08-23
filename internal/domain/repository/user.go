package repository

import (
	"github.com/Ndraaa15/musiku/internal/infrastructure/postgresql"
)

type UserRepository struct {
	db *postgresql.DB
}

type UserRepositoryImpl interface{}

func NewUserRepository(db *postgresql.DB) UserRepositoryImpl {
	return &UserRepository{db}
}
