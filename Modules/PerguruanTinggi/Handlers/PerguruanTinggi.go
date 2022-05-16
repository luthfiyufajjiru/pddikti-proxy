package Handlers

import (
	"PDDiktiProxyAPI/Modules/PerguruanTinggi/DataTransferObjects"
	"PDDiktiProxyAPI/Modules/PerguruanTinggi/Services"
	Services2 "PDDiktiProxyAPI/Modules/Search/Services"
	"PDDiktiProxyAPI/Modules/ServerCaches"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetUniversities() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		results, err := ServerCaches.GetUniversities()

		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		} else if results == nil || *results == nil {
			return fiber.NewError(http.StatusNoContent)
		}

		return ctx.JSON(results)
	}
}

func GetUniversityByKodePt() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		kodePt := ctx.Params("query")
		var emptyUni DataTransferObjects.PerguruanTinggiDTO

		if kodePt == "" {
			return fiber.NewError(http.StatusBadRequest, "kode perguruan tinggi tidak boleh kosong!")
		}

		result, err := Services.GetUniversity(kodePt)

		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		} else if result == emptyUni {
			return fiber.NewError(http.StatusNoContent)
		}

		return ctx.JSON(result)
	}
}

func SearchUniversitiesByName() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		query := ctx.Params("query")
		if query == "" {
			return fiber.NewError(http.StatusBadRequest, "nama perguruan tinggi tidak boleh kosong!")
		}

		results, err := Services2.SearchUniversity(query)

		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		} else if results == nil {
			return fiber.NewError(http.StatusNoContent)
		}

		return ctx.JSON(results)

	}
}
