package router

import (
	"shop/delivery/handler"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func InitToko(router fiber.Router, serv service.Toko) {
	toko := router.Group("/toko")
	toko.Get("/my", handler.GetMyToko(serv))
	toko.Put("/:id", handler.UpdateToko(serv))
	toko.Get("/:id", handler.GetTokoByID(serv))
	toko.Get("/", handler.GetAllToko(serv))
}
