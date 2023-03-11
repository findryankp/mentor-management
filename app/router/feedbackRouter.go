package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"immersiveApp/app/middlewares"
	_feedbackData "immersiveApp/features/feedbacks/data"
	_feedbackHandler "immersiveApp/features/feedbacks/delivery"
	_feedbackService "immersiveApp/features/feedbacks/service"
)

func FeedbackRouter(db *gorm.DB, e *echo.Echo) {
	data := _feedbackData.New(db)
	service := _feedbackService.New(data)
	handler := _feedbackHandler.New(service)

	g := e.Group("/feedbacks")
	g.Use(middlewares.Authentication)
	g.GET("", handler.GetAll)      //not used
	g.GET("/:id", handler.GetById) //not used
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}
