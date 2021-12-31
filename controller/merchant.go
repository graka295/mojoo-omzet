package controller

import (
	"mojoo/omzet/container"
	"mojoo/omzet/services"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type MerchantController interface {
	GetOmzetMerchant(ctx echo.Context) error
}

// MerchantControllerImpl struct
type MerchantControllerImpl struct {
	container container.Container
	services  services.Merchant
}

func MerchantControllerNew(container container.Container, services services.Merchant) MerchantController {
	return &MerchantControllerImpl{
		container: container,
		services:  services,
	}
}

func (c MerchantControllerImpl) GetOmzetMerchant(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	return c.services.GetOmzetMerchant(ctx, int(id))
}
