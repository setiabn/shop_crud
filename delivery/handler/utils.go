package handler

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// func getToken(c *fiber.Ctx) string {}
// func parseIdPathParams(c *fiber.Ctx) string {}

var (
	ErrWrongParams    = errors.New("wrong parameter")
	ErrInternalServer = errors.New("error processing request")
)

func respSuccess(method string, data interface{}) fiber.Map {
	return fiber.Map{
		"data":    data,
		"errors":  nil,
		"message": "Succeed to " + method + " data",
		"status":  true,
	}
}

func respError(method string, errs ...error) fiber.Map {

	var errStrings []string
	for _, e := range errs {
		errStrings = append(errStrings, e.Error())
	}
	return fiber.Map{
		"data":    nil,
		"errors":  errStrings,
		"message": "Failed to " + method + " data",
		"status":  false,
	}
}

func getPathVariableId(c *fiber.Ctx) (uint, error) {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return 0, ErrWrongParams
	}

	if id <= 0 {
		return 0, ErrWrongParams
	}

	return uint(id), nil
}

func parseDate(t string) (time.Time, error) {

	parsedDate, err := time.Parse("02/01/2006", t)
	if err != nil {
		return time.Now(), err
	}

	return parsedDate, nil
}
