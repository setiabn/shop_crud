package router

import (
	"shop/delivery/handler"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func InitCategory(router fiber.Router, serv service.Category) {

	category := router.Group("/category")

	category.Get("/", handler.GetAllCategory(serv))
	category.Get("/:id", handler.GetCategoryByID(serv))
	category.Post("/", handler.CreateCategory(serv))
	category.Put("/:id", handler.UpdateCategoryByID(serv))
	category.Delete("/:id", handler.DeleteCategoryByID(serv))
}
