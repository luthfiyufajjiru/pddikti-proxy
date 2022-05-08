package main

import "github.com/gofiber/fiber/v2"

func RunServer(address string) error {
	app := fiber.New()

	//api := app.Group("/api")

	//v1collection := api.Group("/v1")

	//v1collection.Mount("/v1/users", Account.GetRoutesV1(h))

	return app.Listen(address)
}
