package main

import (
	"github.com/arvinpaundra/go-rent-bike/configs"
	"github.com/arvinpaundra/go-rent-bike/database"
	"github.com/arvinpaundra/go-rent-bike/internal/route"
	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitConfig()
	database.InitMysqlDatabase()

	e := echo.New()

	route.New(database.DB, e)

	e.Logger.Fatal(e.Start(configs.Cfg.AppPort))
}
