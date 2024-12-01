package main

import (
	"femas/belajar-fiber/database"
	"femas/belajar-fiber/database/migrations"
	"femas/belajar-fiber/helpers"
	"femas/belajar-fiber/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public", fiber.Static{
		//Index: "tes.txt",
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return helpers.Success(c, "hello world")
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		return helpers.Err(c, 400, "Awok")
	})

	database.Init()

	migrations.Migration()

	routes.RouteInit(app)

	err := app.Listen(":8090")
	if err != nil {
		fmt.Println(err)
	}

}
