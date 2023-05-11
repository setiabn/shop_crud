package router

import (
	"shop/delivery/handler"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func InitAuth(router fiber.Router, servAuth service.Auth, servProvCity service.ProvCity) {

	auth := router.Group("/auth")
	auth.Post("/register", handler.Register(servAuth))
	auth.Post("/login", handler.Login(servAuth, servProvCity))
}
