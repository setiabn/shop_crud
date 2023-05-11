package model

import "gorm.io/gorm"

type Toko struct {
	gorm.Model
	NamaToko string `gorm:"column:nama_toko;type:varchar(255)" json:"nama_toko" validate:"required"`
	URLFoto  string `gorm:"column:url_foto;type:varchar(255)" json:"url_foto" validate:"required"`

	UserID      uint
	Products    []Product
	DetailTrxs  []DetailTrx
	LogProducts []LogProduct
}
