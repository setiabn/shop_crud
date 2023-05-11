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
		"status":  true,
		"message": "Succeed to " + method + " data",
		"errors":  nil,
		"data":    data,
	}
}

func respError(method string, errs ...error) fiber.Map {

	var errStrings []string
	for _, e := range errs {
		errStrings = append(errStrings, e.Error())
	}
	return fiber.Map{
		"status":  false,
		"message": "Failed to " + method + " data",
		"errors":  errStrings,
		"data":    nil,
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

func parseToUint(numbers ...string) ([]uint, error) {

	var result []uint
	for _, v := range numbers {
		num, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return nil, err
		}
		result = append(result, uint(num))
	}
	return result, nil
}
