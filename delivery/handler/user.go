package handler

import (
	"shop/middleware"
	"shop/model"
	"shop/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetMyProfile(service service.User) fiber.Handler {

	return func(c *fiber.Ctx) error {

		tokenStr := c.Get("Token")
		token, err := middleware.ParseToken(tokenStr)
		if err != nil {
			return middleware.JwtError(c, err)
		}

		user, err := service.Get(token.UserId)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), user))
	}
}

type reqUpdateProfile struct {
	Nama         string `json:"nama"`
	KataSandi    string `json:"kata_sandi"`
	NoTelp       string `json:"no_telp"`
	TanggalLahir string `json:"tanggal_Lahir"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email"`
	IdProvinsi   string `json:"id_provinsi"`
	IdKota       string `json:"id_kota"`
}

func UpdateProfile(service service.User) fiber.Handler {

	return func(c *fiber.Ctx) error {

		tokenStr := c.Get("Token")
		token, err := middleware.ParseToken(tokenStr)
		if err != nil {
			return middleware.JwtError(c, err)
		}

		var updatedProfile reqUpdateProfile
		if err := c.BodyParser(&updatedProfile); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		birthday, err := time.Parse("02/01/2006", updatedProfile.TanggalLahir)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		if err := c.BodyParser(&updatedProfile); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		updatedUser := model.User{
			Nama:         updatedProfile.Nama,
			KataSandi:    updatedProfile.KataSandi,
			NoTelp:       updatedProfile.NoTelp,
			TanggalLahir: birthday,
			Pekerjaan:    updatedProfile.Pekerjaan,
			Email:        updatedProfile.Email,
			IDProvinsi:   updatedProfile.IdProvinsi,
			IDKota:       updatedProfile.IdKota,
		}
		updatedUser.ID = token.UserId

		user, err := service.Update(updatedUser)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), user))

	}
}
