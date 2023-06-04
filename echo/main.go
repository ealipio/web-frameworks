package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New() // echo instance

	// route
	e.GET("/", handler)
	e.GET("/user/:userId", handlerUser)
	fmt.Println("Hello Echo")

	// middleware
	e.Use(middleware.Logger())

	// run app
	e.Start(":8080")
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Echo!")
}

func handlerUser(c echo.Context) error {
	userId := c.Param("userId")
	return c.String(http.StatusOK, userId)
}
