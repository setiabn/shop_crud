package router

import (
	"shop/delivery/handler"
	"shop/middleware"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func initAlamat(router fiber.Router, serv service.Alamat) {

	alamat := router.Group("/alamat", middleware.UserOnly())

	alamat.Get("/", handler.GetMyAlamat(serv))
	alamat.Get("/:id", handler.GetAlamatByID(serv))
	alamat.Post("/", handler.CreateAlamat(serv))
	alamat.Put("/:id", handler.UpdateAlamatByID(serv))
	alamat.Delete("/:id", handler.DeleteAlamatByID(serv))
}
