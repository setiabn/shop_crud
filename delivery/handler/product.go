package handler

import (
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(service service.Product) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}

func GetProductByID(service service.Product) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}

func CreateProduct(service service.Product) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}
func UpdateProductByID(service service.Product) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}

func DeleteProductByID(service service.Product) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}
