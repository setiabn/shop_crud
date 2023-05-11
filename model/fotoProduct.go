package model

type FotoProduct struct {
	BasicModelShowID
	Url string `gorm:"column:url;type:varchar(255)" json:"url" validate:"required"`

	ProductID uint `gorm:"column:product_id;type:int" json:"-"`
}
