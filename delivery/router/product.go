package router

import (
	"shop/delivery/handler"
	"shop/middleware"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func InitProduct(router fiber.Router, serv service.Product) {

	product := router.Group("/product", middleware.UserOnly())

	product.Get("/", handler.GetAllProduct(serv))
	product.Get("/:id", handler.GetProductByID(serv))
	product.Post("/", handler.CreateProduct(serv))
	product.Put("/:id", handler.UpdateProductByID(serv))
	product.Delete("/:id", handler.DeleteProductByID(serv))
}
