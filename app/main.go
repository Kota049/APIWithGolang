package main

import (
	"main/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	e.GET("/", controllers.Index)
	e.POST("/", controllers.StoreArticle)
	e.PUT("/:id", controllers.UpdateArticle)
	e.DELETE("/:id", controllers.DestroyArticle)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}
