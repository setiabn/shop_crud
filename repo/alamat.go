package repo

import (
	"errors"
	"shop/model"

	"gorm.io/gorm"
)

type Alamat interface {
	Create(alamat model.Alamat) (model.Alamat, error)
	Get(alamatid uint) (model.Alamat, error)
	GetByUserID(userid uint) ([]model.Alamat, error)
	Update(alamat model.Alamat) (model.Alamat, error)
	Delete(alamat model.Alamat) error
}

func NewAlamatRepo(db *gorm.DB) Alamat {
	return &repoalamat{
		DB: db,
	}
}

type repoalamat struct {
	DB *gorm.DB
}

func (r *repoalamat) Create(alamat model.Alamat) (model.Alamat, error) {

	result := r.DB.Create(&alamat)
	if result.Error != nil {
		return model.Alamat{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.Alamat{}, errors.New("not created")
	}

	return alamat, nil
}

func (r *repoalamat) Get(id uint) (model.Alamat, error) {
	var alamat model.Alamat
	result := r.DB.First(&alamat, id)
	if result.Error != nil {
		return model.Alamat{}, result.Error
	}

	return alamat, nil
}

func (r *repoalamat) GetByUserID(userid uint) ([]model.Alamat, error) {
	var alamats []model.Alamat
	result := r.DB.Where(&model.Alamat{UserID: userid}).Find(&alamats)
	if result.Error != nil {
		return []model.Alamat{}, result.Error
	}
	return alamats, nil
}

func (r *repoalamat) Update(alamat model.Alamat) (model.Alamat, error) {

	var olddata model.Alamat

	result := r.DB.First(&olddata, alamat.ID)
	if result.Error != nil {
		return model.Alamat{}, result.Error
	}

	alamat.CreatedAt = olddata.CreatedAt

	result = r.DB.Save(&alamat)
	if result.Error != nil {
		return model.Alamat{}, result.Error
	}

	return alamat, nil
}

func (r *repoalamat) Delete(alamat model.Alamat) error {

	result := r.DB.Where(&model.Alamat{UserID: alamat.UserID}).Delete(&model.Alamat{}, alamat.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
