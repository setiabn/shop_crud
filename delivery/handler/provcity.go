package handler

import (
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func GetListProvincies(serv service.ProvCity) fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := serv.GetAllProvincies()
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), result))
	}
}

func GetListCities(serv service.ProvCity) fiber.Handler {
	return func(c *fiber.Ctx) error {

		provId := c.Params("id")

		result, err := serv.GetAllCities(provId)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), result))
	}
}

func GetDetaiProvince(serv service.ProvCity) fiber.Handler {
	return func(c *fiber.Ctx) error {
		provId := c.Params("id")

		result, err := serv.GetDetaiProvince(provId)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), result))
	}
}

func GetDetailCity(serv service.ProvCity) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cityId := c.Params("id")

		result, err := serv.GetDetailCity(cityId)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), result))
	}
}
