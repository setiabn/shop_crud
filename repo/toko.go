package repo

import (
	"shop/model"

	"gorm.io/gorm"
)

type Toko interface {
	Get(tokoid uint) (model.Toko, error)
	GetAll(limit uint, page uint) ([]model.Toko, error)
	Update(toko model.Toko) (model.Toko, error)
	Delete(tokoid uint) error
}

func NewTokoRepo(db *gorm.DB) Toko {
	return &repotoko{
		DB: db,
	}
}

type repotoko struct {
	DB *gorm.DB
}

func (r *repotoko) Get(id uint) (model.Toko, error) {
	var toko model.Toko
	result := r.DB.First(&toko, id)
	if result.Error != nil {
		return model.Toko{}, result.Error
	}

	return toko, nil
}

func (r *repotoko) GetAll(limit uint, page uint) ([]model.Toko, error) {
	var tokos []model.Toko

	offset := int(limit) * int(page-1)
	result := r.DB.Limit(int(limit)).Offset(offset)
	if result.Error != nil {
		return []model.Toko{}, result.Error
	}
	return tokos, nil
}

func (r *repotoko) Update(toko model.Toko) (model.Toko, error) {

	var olddata model.Toko

	result := r.DB.First(&olddata, toko.ID)
	if result.Error != nil {
		return model.Toko{}, result.Error
	}

	toko.CreatedAt = olddata.CreatedAt

	result = r.DB.Create(&toko)
	if result.Error != nil {
		return model.Toko{}, result.Error
	}

	return toko, nil
}

func (r *repotoko) Delete(id uint) error {

	result := r.DB.Delete(&model.Toko{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
