package app

import (
	"mojoo/omzet/container"
	"mojoo/omzet/controller"
	"mojoo/omzet/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

// NewRouter routing file
func NewRouter(e *echo.Echo, container container.Container) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	auth := controller.AuthControllerNew(container, services.NewAuthServices(container))
	api := e.Group("/api")
	api.POST("/login", auth.Login)
	v1 := api.Group("/admin", IsLoggedIn)
	merchant := controller.MerchantControllerNew(container, services.NewMerchant(container))
	v1.POST("/merchant", merchant.GetOmzetMerchant)
	outlets := controller.OutletsControllerNew(container, services.NewOustLets(container))
	v1.POST("/outlets", outlets.GetOmzetOutlets)
}
