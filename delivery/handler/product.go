package handler

import (
	"fmt"
	"shop/middleware"
	"shop/model"
	"shop/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllProduct(service service.Product) fiber.Handler {

	return func(c *fiber.Ctx) error {

		// namaProduk := c.FormValue("nama_produk")
		limit := c.FormValue("limit")
		page := c.FormValue("page")
		categoryId := c.FormValue("category_id")
		tokoId := c.FormValue("toko_id")
		maxHarga := c.FormValue("max_harga")
		minHarga := c.FormValue("min_harga")

		numbers, err := parseToUint(limit, page, categoryId, tokoId, maxHarga, minHarga)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		products, err := service.GetAll(numbers[0], numbers[1], numbers[2], numbers[3])
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), fiber.Map{
			"data":  products,
			"page":  page,
			"limit": limit,
		}))
	}
}

func GetProductByID(service service.Product) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}

func CreateProduct(service service.Product) fiber.Handler {

	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Token")
		token, err := middleware.ParseToken(tokenStr)
		if err != nil {
			return middleware.JwtError(c, err)
		}

		namaProduk := c.FormValue("nama_produk")
		hargaReseller := c.FormValue("harga_reseller")
		hargaKonsumen := c.FormValue("harga_konsumen")
		deskripsi := c.FormValue("deskripsi")

		categoryId, err := strconv.Atoi(c.FormValue("category_id"))
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}
		stok, err := strconv.Atoi(c.FormValue("stok"))
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}

		form, err := c.MultipartForm()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(respError(c.Method(), err))
		}
		// Get all files  and upload
		files := form.File["photos"]
		var fotoProducts []model.FotoProduct
		for _, file := range files {

			id := uuid.New().String()
			filename := fmt.Sprintf("public/assets/produk/%v-%s", id, file.Filename)
			err := c.SaveFile(file, filename)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(respError(c.Method(), err))
			}
			fotoProducts = append(fotoProducts, model.FotoProduct{
				Url:       filename,
				ProductID: token.UserId,
			})
		}

		newProduct := model.Product{
			NamaProduct:   namaProduk,
			HargaReseller: hargaReseller,
			HargaKonsumen: hargaKonsumen,
			Stok:          uint(stok),
			Deskripsi:     deskripsi,
			TokoID:        token.TokoID,
			FotoProducts:  fotoProducts,
			CategoryID:    uint(categoryId),
		}

		product, err := service.Create(newProduct)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(respError(c.Method(), err))
		}

		return c.JSON(respSuccess(c.Method(), product))
	}
}
func UpdateProductByID(service service.Product) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}

func DeleteProductByID(service service.Product) fiber.Handler {

	return func(c *fiber.Ctx) error {

		return c.JSON("Ok")
	}
}
