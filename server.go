package main

import (
	"PDDiktiProxyAPI/Modules/General"
	"github.com/gofiber/fiber/v2"
)

func RunServer(address string) error {
	app := fiber.New()

	api := app.Group("/api")

	v1collection := api.Group("/v1")
	v1collection.Mount("/perguruan-tinggi", General.GetRoutesV1())

	//app.Server().GetOpenConnectionsCount()

	//location, _ := time.LoadLocation("Asia/Jakarta")
	//x := gocron.NewScheduler(location)
	//x.Every(1).Day().At("05:51").Do(ServerCaches.CleanUniversities)
	//x.StartAsync()

	return app.Listen(address)
}
