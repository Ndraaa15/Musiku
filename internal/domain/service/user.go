package service

import "github.com/Ndraaa15/musiku/internal/domain/repository"

type UserService struct {
	ur *repository.UserRepositoryImpl
}

type UserServiceImpl interface {
}

func NewUserService(ur repository.UserRepositoryImpl) UserServiceImpl {
	return &UserService{}
}
