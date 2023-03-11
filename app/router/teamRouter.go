package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"immersiveApp/app/middlewares"
	_teamData "immersiveApp/features/teams/data"
	_teamHandler "immersiveApp/features/teams/delivery"
	_teamService "immersiveApp/features/teams/service"
)

func TeamRouter(db *gorm.DB, e *echo.Echo) {
	data := _teamData.New(db)
	service := _teamService.New(data)
	handler := _teamHandler.New(service)

	g := e.Group("/teams")
	g.Use(middlewares.Authentication)
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}
