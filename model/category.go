package model

type Category struct {
	BasicModelShowID
	NamaCategory string `gorm:"column:nama_category;type:varchar(255)" json:"nama_category" validate:"required"`
}
