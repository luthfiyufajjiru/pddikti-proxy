package General

import (
	"PDDiktiProxyAPI/Modules/General/Handlers"
	"github.com/gofiber/fiber/v2"
)

func GetRoutesV1() (routes *fiber.App) {
	routes = fiber.New()
	routes.Get("", Handlers.GetUniversities())
	routes.Get("/search/:query", Handlers.SearchUniversitiesByName())
	routes.Get("/:mode/:query", Handlers.GetUniversity())
	routes.Get("/k/:query/daftar-prodi", Handlers.GetProdi())
	return
}
