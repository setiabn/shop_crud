package router

import (
	"shop/delivery/handler"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func InitProvCity(router fiber.Router, serv service.ProvCity) {

	provcity := router.Group("/provcity")
	provcity.Get("/listprovincies", handler.GetListProvincies(serv))
	provcity.Get("/listcities/:id", handler.GetListCities(serv))
	provcity.Get("/detailprovince/:id", handler.GetDetaiProvince(serv))
	provcity.Get("/detailcity/:id", handler.GetDetailCity(serv))

}
