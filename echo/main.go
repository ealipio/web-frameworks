package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// echo instance
	e := echo.New()

	// routes
	registerRoutes(e)

	// middleware
	e.Use(middleware.Logger())

	// run app
	e.Start(":8080")
}
