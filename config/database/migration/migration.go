package migration

import (
	"shop/model"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Toko{},
		&model.Trx{},
		&model.Alamat{},
		&model.DetailTrx{},
		&model.Product{},
		&model.LogProduct{},

		&model.Category{},
		&model.FotoProduct{},
	)
	if err != nil {
		panic(err)
	}
}
