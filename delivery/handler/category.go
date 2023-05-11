package handler

import (
	"shop/model"
	"shop/service"

	"github.com/gofiber/fiber/v2"
)

func respCategory(data model.Category) fiber.Map {
	return fiber.Map{
		"id":            data.ID,
		"nama_category": data.NamaCategory,
	}
}

func GetAllCategory(service service.Category) fiber.Handler {

	return func(c *fiber.Ctx) error {

		result, err := service.GetAll()
		if err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(respError(c.Method(), err))
		}

		var arr []fiber.Map
		for _, category := range result {
			arr = append(arr, respCategory(category))
		}

		return c.JSON(respSuccess(c.Method(), arr))
	}
}

func GetCategoryByID(service service.Category) fiber.Handler {

	return func(c *fiber.Ctx) error {
		id, err := getPathVariableId(c)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}
		result, err := service.Get(id)
		if err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), respCategory(result)))
	}
}

func CreateCategory(service service.Category) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var category model.Category

		if err := c.BodyParser(&category); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), ErrWrongParams))
		}

		result, err := service.Create(category)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		c.Status(fiber.StatusCreated)
		return c.JSON(respSuccess(c.Method(), respCategory(result)))
	}
}

func UpdateCategoryByID(service service.Category) fiber.Handler {

	return func(c *fiber.Ctx) error {

		// Get the id
		id, err := getPathVariableId(c)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		// Get category data
		var category model.Category

		if err := c.BodyParser(&category); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), ErrWrongParams))
		}

		category.ID = id

		result, err := service.Update(category)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), respCategory(result)))
	}
}

func DeleteCategoryByID(service service.Category) fiber.Handler {

	return func(c *fiber.Ctx) error {

		// Get the id
		id, err := getPathVariableId(c)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		if err := service.Delete(id); err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), ""))
	}
}
