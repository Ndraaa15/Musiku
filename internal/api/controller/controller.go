package controller

import (
	"github.com/Ndraaa15/musiku/internal/domain/service"
)

type Controller struct {
	us service.UserServiceImpl
}

func NewController(us service.UserServiceImpl) *Controller {
	return &Controller{
		us: us,
	}
}
