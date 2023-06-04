package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()
	// route
	app.Get("/", handler)
	// run
	app.Listen(":8080")

}

// handlers
func handler(c *fiber.Ctx) error {
	return c.SendString("Hello, Fiber 👋!")
}
