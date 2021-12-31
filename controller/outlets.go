package controller

import (
	"mojoo/omzet/container"
	"mojoo/omzet/services"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type OutletsController interface {
	GetOmzetOutlets(ctx echo.Context) error
}

// OutletsControllerImpl struct
type OutletsControllerImpl struct {
	container container.Container
	services  services.OutLets
}

func OutletsControllerNew(container container.Container, services services.OutLets) OutletsController {
	return &OutletsControllerImpl{
		container: container,
		services:  services,
	}
}

func (c OutletsControllerImpl) GetOmzetOutlets(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	return c.services.GetOmszetOutLets(ctx, int(id))
}
