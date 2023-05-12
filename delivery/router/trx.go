package router

import (
	"shop/delivery/handler"
	"shop/middleware"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func InitTrx(router fiber.Router, serv service.Trx) {

	trx := router.Group("/trx", middleware.UserOnly())

	trx.Get("/", handler.GetAllMyTrx(serv))
	trx.Get("/:id", handler.GetTrxByID(serv))
	trx.Post("/", handler.CreateTrx(serv))
}
