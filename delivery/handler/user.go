package handler

import (
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func GetMyProfile(service service.User) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}

func UpdateProfile(service service.User) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}
