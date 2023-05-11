package model

type DetailTrx struct {
	BasicModelShowID

	Kuantitas  int `gorm:"column:kuantitas;type:int" validate:"required"`
	HargaTotal int `gorm:"column:harga_total;type:int" validate:"required"`

	TrxID  uint
	TokoID uint

	LogProducts []LogProduct
}
