package PerguruanTinggi

import (
	"PDDiktiProxyAPI/Modules/PerguruanTinggi/Handlers"
	"github.com/gofiber/fiber/v2"
)

func GetRoutesV1() (routes *fiber.App) {
	routes = fiber.New()
	routes.Get("", Handlers.GetUniversities())
	routes.Get("/:query", Handlers.GetUniversityByKodePt())
	routes.Get("search/:query", Handlers.SearchUniversitiesByName())
	return
}
