package repository

import (
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
)

type UserRepository struct {
	db *mysql.DB
}

type UserRepositoryImpl interface{}

func NewUserRepository(db *mysql.DB) UserRepositoryImpl {
	return &UserRepository{db}
}
