package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_userData "clean-arch/features/users/data"
	_userHandler "clean-arch/features/users/delivery"
	_userService "clean-arch/features/users/service"
)

func UserRouter(db *gorm.DB, e *echo.Echo) {
	data := _userData.New(db)
	service := _userService.New(data)
	handler := _userHandler.New(service)

	g := e.Group("/users")
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Create)
	g.DELETE("/:id", handler.Create)
}
