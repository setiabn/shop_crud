package repo

import (
	"errors"
	"shop/model"

	"github.com/gofiber/fiber/v2"
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

	// Protect from other user
	for _, alamat := range alamats {
		if alamat.UserID != userid {
			return []model.Alamat{}, fiber.ErrUnauthorized
		}
	}
	return alamats, nil
}

func (r *repoalamat) Update(alamat model.Alamat) (model.Alamat, error) {

	var olddata model.Alamat

	result := r.DB.First(&olddata, alamat.ID)
	if result.Error != nil {
		return model.Alamat{}, result.Error
	}

	if olddata.UserID != alamat.UserID {
		return model.Alamat{}, fiber.ErrUnauthorized
	}

	alamat.CreatedAt = olddata.CreatedAt

	result = r.DB.Save(&alamat)
	if result.Error != nil {
		return model.Alamat{}, result.Error
	}

	return alamat, nil
}

func (r *repoalamat) Delete(alamat model.Alamat) error {

	var olddata model.Alamat
	if err := r.DB.First(&olddata).Error; err != nil {
		return err
	}
	if olddata.UserID != alamat.UserID {
		return fiber.ErrUnauthorized
	}

	result := r.DB.Where(&model.Alamat{UserID: alamat.UserID}).Delete(&model.Alamat{}, alamat.ID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errRecordNotFound
	}

	return nil
}
