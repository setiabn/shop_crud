package handler

import (
	"fmt"
	"shop/middleware"
	"shop/model"
	"shop/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllToko(service service.Toko) fiber.Handler {

	return func(c *fiber.Ctx) error {

		query := struct {
			Limit uint `json:"limit"`
			Page  uint `json:"page"`
			// Nama  string `json:"nama"`
		}{}

		if err := c.QueryParser(&query); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		tokos, err := service.GetAll(query.Limit, query.Page)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), tokos))
	}
}

func GetTokoByID(service service.Toko) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := getPathVariableId(c)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		toko, err := service.GetByID(id)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), toko))
	}
}

func GetMyToko(service service.Toko) fiber.Handler {

	return func(c *fiber.Ctx) error {

		tokenStr := c.Get("Token")
		token, err := middleware.ParseToken(tokenStr)
		if err != nil {
			return middleware.JwtError(c, err)
		}

		alamats, err := service.GetByUserID(token.UserId)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), alamats))
	}
}

func UpdateToko(service service.Toko) fiber.Handler {

	return func(c *fiber.Ctx) error {

		tokenStr := c.Get("Token")
		token, err := middleware.ParseToken(tokenStr)
		if err != nil {
			return middleware.JwtError(c, err)
		}

		dataToko := struct {
			NamaToko string `json:"nama_toko"`
		}{}

		if err := c.BodyParser(&dataToko); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		photo, err := c.FormFile("photo")
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		id := uuid.New().String()
		photofilename := fmt.Sprintf("public/assets/toko/%v-%s", id, photo.Filename)

		newToko := model.Toko{
			NamaToko: dataToko.NamaToko,
			UserID:   token.UserId,
			URLFoto:  photofilename,
		}
		newToko.ID = token.UserId

		// Simpan data
		toko, err := service.Update(newToko)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		// Simpan foto
		if err := c.SaveFile(photo, photofilename); err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), toko))
	}
}
