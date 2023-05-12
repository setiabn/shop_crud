package handler

import (
	"shop/middleware"
	"shop/model"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

type registerReqData struct {
	Nama         string `json:"nama"`
	KataSandi    string `json:"kata_sandi"`
	NoTelp       string `json:"no_telp"`
	TanggalLahir string `json:"tanggal_lahir"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email"`
	IdProvinsi   string `json:"id_provinsi"`
	IdKota       string `json:"id_kota"`
}

func Register(service service.Auth) fiber.Handler {

	return func(c *fiber.Ctx) error {

		var registerData registerReqData
		if err := c.BodyParser(&registerData); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		birthday, err := parseDate(registerData.TanggalLahir)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		user := model.User{
			Nama:         registerData.Nama,
			KataSandi:    registerData.KataSandi,
			NoTelp:       registerData.NoTelp,
			TanggalLahir: birthday,
			JenisKelamin: "laki-laki",
			Tentang:      "-",
			Pekerjaan:    registerData.Pekerjaan,
			Email:        registerData.Email,
			IDProvinsi:   registerData.IdProvinsi,
			IDKota:       registerData.IdKota,
			IsAdmin:      false,
		}

		if err := service.Register(user); err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), "Register Succeed"))
	}
}

type loginReqData struct {
	KataSandi string `json:"kata_sandi"`
	NoTelp    string `json:"no_telp"`
}

func loginResp(user model.User, province model.Province, city model.City, token string) fiber.Map {
	return fiber.Map{
		"nama":          user.Nama,
		"no_telp":       user.NoTelp,
		"tanggal_Lahir": user.TanggalLahir.Format("02/01/2006"),
		"tentang":       user.Tentang,
		"pekerjaan":     user.Pekerjaan,
		"email":         user.Email,
		"id_provinsi":   province,
		"id_kota":       city,
		"token":         token,
	}
}

func Login(servAuth service.Auth, servProvCity service.ProvCity) fiber.Handler {

	return func(c *fiber.Ctx) error {

		var loginData loginReqData
		if err := c.BodyParser(&loginData); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		user := model.User{
			NoTelp:    loginData.NoTelp,
			KataSandi: loginData.KataSandi,
		}

		// Cek apakah user bisa login
		user, err := servAuth.Login(user)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		// Generate token from user
		tokenString, err := middleware.GenerateJwtToken(user)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		// Melengkapi informasi mengenai provinsi & kota ----------------------------------
		province, err := servProvCity.GetDetaiProvince(user.IDProvinsi)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		city, err := servProvCity.GetDetailCity(user.IDKota)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}
		// --------------------------------------------------------------------------------

		return c.JSON(respSuccess(
			c.Method(),
			loginResp(user, province, city, tokenString),
		))
	}
}
