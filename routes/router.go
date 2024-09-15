package routes

import (
	"github.com/jamesguru/dev-api/controllers"
	"github.com/labstack/echo/v4"
)

func Setup() {
	e := echo.New()
	e.POST("/books", controllers.AddBooks)
	e.GET("/book/:id", controllers.GetBook)
	e.Start(":3001")
}
