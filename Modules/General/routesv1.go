package General

import (
	"PDDiktiProxyAPI/Modules/General/Handlers"
	"github.com/gofiber/fiber/v2"
)

func GetRoutesV1() (routes *fiber.App) {
	routes = fiber.New()
	routes.Get("/universities", Handlers.GetUniversities())
	routes.Get("/university/search/:query", Handlers.SearchUniversitiesByName())
	routes.Get("/university/:mode/:query", Handlers.GetUniversity())
	routes.Get("/university/c/:query/majors", Handlers.GetProdi())
	return
}
