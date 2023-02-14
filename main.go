package main

import "github.com/gofiber/fiber/v2"
import "fiber-project-example/routes"

func main() {
	app := fiber.New()

	routes.UseBooksRoute(app)

	app.Listen(":4000")
}