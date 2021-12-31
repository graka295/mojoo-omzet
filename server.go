package main

import (
	"mojoo/omzet/app"
	"mojoo/omzet/container"
	"mojoo/omzet/helper"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	e := echo.New()
	configImpl := helper.Load()
	log := helper.NewLog()
	db := helper.NewDB(configImpl)

	if configImpl.GetConfig().Database.Migration {
		db.Migrate()
	}
	response := helper.NewResponse()
	validate := helper.NewValidate(validator.New())
	container := container.NewContainer(db, configImpl, log, response, validate)
	app.NewRouter(e, container)

	e.Logger.Fatal(e.Start(":1323"))
}
