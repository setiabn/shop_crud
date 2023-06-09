package handler

import (
	"shop/middleware"
	"shop/model"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func GetMyAlamat(service service.Alamat) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token, err := getToken(c)
		if err != nil {
			return middleware.JwtError(c, err)
		}

		// userid dari jwt, sehingga aman dari user lain
		alamats, err := service.GetByUserID(token.UserId)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}
		// Proteksi dari user lain
		if err := isOwner(c, alamats[0].UserID); err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(respError(c.Method(), fiber.ErrUnauthorized))
		}

		return c.JSON(respSuccess(c.Method(), alamats))
	}
}

func GetAlamatByID(service service.Alamat) fiber.Handler {

	return func(c *fiber.Ctx) error {

		id, err := getPathVariableId(c)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		alamat, err := service.GetByID(id)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		// Proteksi dari user lain
		if err := isOwner(c, alamat.UserID); err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(respError(c.Method(), fiber.ErrUnauthorized))
		}

		return c.JSON(respSuccess(c.Method(), alamat))
	}
}

func CreateAlamat(service service.Alamat) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var newAlamat model.Alamat
		err := c.BodyParser(&newAlamat)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		token, err := getToken(c)
		if err != nil {
			return middleware.JwtError(c, err)
		}
		newAlamat.UserID = token.UserId

		alamat, err := service.Create(newAlamat)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(respSuccess(c.Method(), alamat))
	}
}

func UpdateAlamatByID(service service.Alamat) fiber.Handler {

	return func(c *fiber.Ctx) error {

		alamatid, err := getPathVariableId(c)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		var newAlamat model.Alamat
		err = c.BodyParser(&newAlamat)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		token, err := getToken(c)
		if err != nil {
			return middleware.JwtError(c, err)
		}
		newAlamat.UserID = token.UserId
		newAlamat.ID = alamatid

		alamat, err := service.Update(newAlamat)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(respSuccess(c.Method(), alamat))
	}
}

func DeleteAlamatByID(service service.Alamat) fiber.Handler {

	return func(c *fiber.Ctx) error {

		alamatid, err := getPathVariableId(c)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		var deletedAlamat model.Alamat

		token, err := getToken(c)
		if err != nil {
			return middleware.JwtError(c, err)
		}
		deletedAlamat.UserID = token.UserId
		deletedAlamat.ID = alamatid

		err = service.Delete(deletedAlamat)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(respSuccess(c.Method(), ""))
	}
}
