package handler

import (
	"shop/middleware"
	"shop/model"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func GetAllMyTrx(service service.Trx) fiber.Handler {

	return func(c *fiber.Ctx) error {

		tokenStr := c.Get("Token")
		token, err := middleware.ParseToken(tokenStr)
		if err != nil {
			return middleware.JwtError(c, err)

		}
		trxs, err := service.GetByUserID(token.UserId)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), trxs))
	}
}

func GetTrxByID(service service.Trx) fiber.Handler {

	return func(c *fiber.Ctx) error {

		tokenStr := c.Get("Token")
		token, err := middleware.ParseToken(tokenStr)
		if err != nil {
			return middleware.JwtError(c, err)

		}
		trx, err := service.GetByID(token.UserId)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), trx))
	}
}

func CreateTrx(service service.Trx) fiber.Handler {

	return func(c *fiber.Ctx) error {
		// token := middleware.GetToken(c)

		bodyData := struct {
			MethodBayar string
			AlamatKirim uint
			DetailTrx   []model.DetailTrx
		}{}
		c.BodyParser(&bodyData)

		// TODO belum ...
		// newTrx := model.Trx{

		// }

		// trx, err := service.Create(token.UserId)
		// if err != nil {
		// 	c.Status(fiber.StatusUnauthorized)
		// 	return c.JSON(respError(c.Method(), err))
		// }

		return c.JSON(respSuccess(c.Method(), "Belum"))
	}
}
