package Handlers

import (
	"PDDiktiProxyAPI/Modules/General/DataTransferObjects"
	"PDDiktiProxyAPI/Modules/General/Services"
	Services2 "PDDiktiProxyAPI/Modules/Search/Services"
	"PDDiktiProxyAPI/Modules/ServerCaches"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
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

func GetUniversity() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		query := ctx.Params("query")
		mode := ctx.Params("mode")

		if x := strings.ToLower(mode); x != "k" && x != "n" {
			return fiber.NewError(http.StatusNotFound)
		}

		var emptyUni DataTransferObjects.PerguruanTinggiDTO

		if query == "" {
			return fiber.NewError(http.StatusBadRequest, "kode perguruan tinggi tidak boleh kosong!")
		}

		var (
			result DataTransferObjects.PerguruanTinggiDTO
			err    error
		)

		if x := strings.ToLower(mode); x == "k" {
			result, err = Services.GetUniversityByKode(query)
		} else if x == "n" {
			result, err = Services.GetUniversityByName(query)
		}

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

func GetProdi() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		query := ctx.Params("query")
		_pt, err := Services.GetUniversityByKode(query)
		query = strings.ReplaceAll(_pt.IdSp, " ", "")
		result := ServerCaches.GetProdi(query)
		if err != nil {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		} else if result == nil {
			return fiber.NewError(http.StatusNoContent)
		}
		return ctx.JSON(result)
	}
}
