package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"immersiveApp/app/middlewares"
	_classData "immersiveApp/features/classes/data"
	_classHandler "immersiveApp/features/classes/delivery"
	_classService "immersiveApp/features/classes/service"
)

func ClassRouter(db *gorm.DB, e *echo.Echo) {
	data := _classData.New(db)
	service := _classService.New(data)
	handler := _classHandler.New(service)

	g := e.Group("/classes")
	g.Use(middlewares.Authentication)
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}
