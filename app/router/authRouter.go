package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"clean-arch/app/middlewares"
	_authData "clean-arch/features/auth/data"
	_authHandler "clean-arch/features/auth/delivery"
	_authService "clean-arch/features/auth/service"
)

func AuthRouter(db *gorm.DB, e *echo.Echo) {
	data := _authData.New(db)
	service := _authService.New(data)
	handler := _authHandler.New(service)

	g := e.Group("/auth")
	g.POST("/register", handler.Register)
	g.POST("/login", handler.Login)

	g.Use(middlewares.Authentication)
	g.GET("/users", handler.GetUserLogin)

	// g.POST("/change-password", handler.Create)
	// g.POST("/forget-password", handler.Create)
}
