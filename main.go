package main

import (
	"clean-arch/app/configs"
	"clean-arch/app/database"
	"clean-arch/app/router"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	db := database.InitDBMysql(*cfg)
	database.InitMigration(db)

	e := echo.New()
	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}
