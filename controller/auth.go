package controller

import (
	"mojoo/omzet/container"
	"mojoo/omzet/services"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Login(ctx echo.Context) error
}

// AuthControllerImpl struct
type AuthControllerImpl struct {
	container container.Container
	services  services.AuthServices
}

func AuthControllerNew(container container.Container, services services.AuthServices) AuthController {
	return AuthControllerImpl{
		container: container,
		services:  services,
	}
}

func (c AuthControllerImpl) Login(ctx echo.Context) error {
	return c.services.Login(ctx)
}
