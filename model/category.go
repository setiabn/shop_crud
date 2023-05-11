package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	NamaCategory string `gorm:"column:nama_category;type:varchar(255)" json:"nama_category" validate:"required"`
}
