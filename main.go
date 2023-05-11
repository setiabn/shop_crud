package main

import (
	"shop/config"
	"shop/config/database"
	"shop/delivery/router"
	"shop/model"
	"shop/repo"
	"shop/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

func mockUser() model.User {
	return model.User{
		Nama:         "santi",
		KataSandi:    "kkkaaaca",
		NoTelp:       "0866323",
		TanggalLahir: time.Date(2006, 1, 1, 1, 1, 1, 1, &time.Location{}),
		JenisKelamin: "laki",
		Tentang:      "tidak ada",
		Pekerjaan:    "kode",
		Email:        "mail@gmail.com",
		IDProvinsi:   "1231",
		IDKota:       "adsa",
		IsAdmin:      false,
	}
}

func main() {
	db := database.ConnectDB()

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	repoAlamat := repo.NewAlamatRepo(db)
	repoCategory := repo.NewCategoryRepo(db)
	repoDetailTrx := repo.NewDetailTrxRepo(db)
	repoFotoProduct := repo.NewFotoProductRepo(db)
	repoLogProduct := repo.NewLogProductRepo(db)
	repoProduct := repo.NewProductRepo(db)
	repoToko := repo.NewTokoRepo(db)
	repoTrx := repo.NewTrxRepo(db)
	repoUser := repo.NewUserRepo(db)
	repoProvCity := repo.NewProvCityRepo()

	servAuth := service.NewServiceAuth(repoUser, repoToko, repoTrx, repoAlamat, repoProvCity)
	servCategory := service.NewServiceCategory(repoCategory)
	servProduct := service.NewServiceProduct(repoProduct, repoLogProduct, repoFotoProduct, repoCategory)
	servToko := service.NewServiceToko(repoToko)
	servTrx := service.NewServiceTrx(repoTrx, repoDetailTrx)
	servUser := service.NewServiceUser(repoUser, repoToko, repoTrx, repoAlamat, repoProvCity)
	servAlamat := service.NewServiceAlamat(repoAlamat)
	servProvCity := service.NewServiceProvCity(repoProvCity)

	// create admin if not exist
	admin := model.User{
		Nama:         "admin",
		KataSandi:    "password",
		NoTelp:       "0812345678",
		TanggalLahir: time.Date(1994, 3, 12, 0, 0, 0, 0, time.UTC),
		JenisKelamin: "pria",
		Tentang:      "Admin",
		Pekerjaan:    "admin",
		Email:        "admin@min.com",
		IDProvinsi:   "11",
		IDKota:       "1101",
		IsAdmin:      true,
	}
	servAuth.Register(admin)

	router.InitAuth(v1, servAuth, servProvCity)
	router.InitCategory(v1, servCategory)
	router.InitProduct(v1, servProduct)
	router.InitProvCity(v1, servProvCity)
	router.InitToko(v1, servToko)
	router.InitTrx(v1, servTrx)
	router.InitUser(v1, servUser, servAlamat)

	v1.Static("/public", "/public")

	app.Listen(config.Get("HOST"))
}
