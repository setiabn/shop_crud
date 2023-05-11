package model

type LogProduct struct {
	BasicModel

	NamaProduct   string `gorm:"column:nama_product;type:varchar(255)" json:"nama_product" validate:"required"`
	Slug          string `gorm:"column:slug;type:varchar(255)" json:"slug" validate:"required"`
	HargaReseller string `gorm:"column:harga_reseller;type:varchar(255)" json:"harga_reseller" validate:"required"`
	HargaKonsumen string `gorm:"column:harga_konsumen;type:varchar(255)" json:"harga_konsumen" validate:"required"`
	Deskripsi     string `gorm:"column:deskripsi;type:text" json:"deskripsi" validate:"required"`

	ProductID uint

	TokoID      uint
	DetailTrxID uint

	CategoryID uint
	Category   Category
}
