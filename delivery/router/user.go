package router

import (
	"shop/delivery/handler"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func InitUser(router fiber.Router, servUser service.User, servAlamat service.Alamat) {
	user := router.Group("/user")
	user.Get("/", handler.GetMyProfile(servUser))
	user.Put("/", handler.UpdateProfile(servUser))

	initAlamat(user, servAlamat)
}
