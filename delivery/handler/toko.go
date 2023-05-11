package handler

import (
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func GetAllToko(service service.Toko) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}

func GetTokoByID(service service.Toko) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}

func GetMyToko(service service.Toko) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}
func UpdateToko(service service.Toko) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}
