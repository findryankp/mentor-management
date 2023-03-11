package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"immersiveApp/app/middlewares"
	_statusData "immersiveApp/features/statuses/data"
	_statusHandler "immersiveApp/features/statuses/delivery"
	_statusService "immersiveApp/features/statuses/service"
)

func StatusRouter(db *gorm.DB, e *echo.Echo) {
	data := _statusData.New(db)
	service := _statusService.New(data)
	handler := _statusHandler.New(service)

	g := e.Group("/status")
	g.Use(middlewares.Authentication)
	g.GET("", handler.GetAll)

}
