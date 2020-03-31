package main

import "github.com/gofiber/fiber"

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, from fiber")
	})

	app.Listen(3000)
}

