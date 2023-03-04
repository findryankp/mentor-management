package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_teamData "clean-arch/features/teams/data"
	_teamHandler "clean-arch/features/teams/delivery"
	_teamService "clean-arch/features/teams/service"
)

func TeamRouter(db *gorm.DB, e *echo.Echo) {
	data := _teamData.New(db)
	service := _teamService.New(data)
	handler := _teamHandler.New(service)

	g := e.Group("/teams")
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Create)
	g.DELETE("/:id", handler.Create)
}
