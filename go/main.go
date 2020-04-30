package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"templates/go/controllers"
)

func main() {
	router := echo.New()

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	router.GET("", controllers.GetIndex)
	router.GET("/", controllers.GetIndex)

	router.Logger.Fatal(router.Start(":8080"))
}
