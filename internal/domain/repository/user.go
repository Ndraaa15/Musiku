package repository

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryImpl interface{}

func NewUserRepository(db *gorm.DB) UserRepositoryImpl {
	return &UserRepository{db}
}
