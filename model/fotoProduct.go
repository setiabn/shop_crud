package model

import "gorm.io/gorm"

type FotoProduct struct {
	gorm.Model
	Url string `gorm:"column:url;type:varchar(255)" json:"url" validate:"required"`

	ProductID uint `gorm:"column:product_id;type:int" json:"-"`
}
