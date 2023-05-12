package router

import (
	"shop/delivery/handler"
	"shop/middleware"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func InitCategory(router fiber.Router, serv service.Category) {

	category := router.Group("/category", middleware.UserOnly())

	category.Get("/", handler.GetAllCategory(serv))
	category.Get("/:id", handler.GetCategoryByID(serv))

	category.Post("/", middleware.AdminOnly(), handler.CreateCategory(serv))
	category.Put("/:id", middleware.AdminOnly(), handler.UpdateCategoryByID(serv))
	category.Delete("/:id", middleware.AdminOnly(), handler.DeleteCategoryByID(serv))
}
