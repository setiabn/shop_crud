package repo

import (
	"errors"
	"shop/model"

	"gorm.io/gorm"
)

type Trx interface {
	Create(trx model.Trx) (model.Trx, error)
	Get(trxid uint) (model.Trx, error)
	GetByUserID(userid uint) ([]model.Trx, error)
	GetAll() ([]model.Trx, error)
	Update(trx model.Trx) (model.Trx, error)
	Delete(trxid uint) error
}

func NewTrxRepo(db *gorm.DB) Trx {
	return &repotrx{
		DB: db,
	}
}

type repotrx struct {
	DB *gorm.DB
}

func (r *repotrx) Create(trx model.Trx) (model.Trx, error) {

	result := r.DB.Create(&trx)
	if result.Error != nil {
		return model.Trx{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.Trx{}, errors.New("not created")
	}

	return trx, nil
}

func (r *repotrx) Get(id uint) (model.Trx, error) {
	var trx model.Trx
	result := r.DB.First(&trx, id)
	if result.Error != nil {
		return model.Trx{}, result.Error
	}

	return trx, nil
}

func (r *repotrx) GetByUserID(userid uint) ([]model.Trx, error) {
	var trx []model.Trx
	result := r.DB.Where(&model.Trx{UserID: userid}).Find(&trx)
	if result.Error != nil {
		return []model.Trx{}, result.Error
	}

	return trx, nil
}

func (r *repotrx) GetAll() ([]model.Trx, error) {

	var categories []model.Trx
	result := r.DB.Find(&categories)
	if result.Error != nil {
		return []model.Trx{}, result.Error
	}

	return categories, nil
}

func (r *repotrx) Update(trx model.Trx) (model.Trx, error) {

	var olddata model.Trx

	result := r.DB.First(&olddata, trx.ID)
	if result.Error != nil {
		return model.Trx{}, result.Error
	}

	trx.CreatedAt = olddata.CreatedAt

	result = r.DB.Create(&trx)
	if result.Error != nil {
		return model.Trx{}, result.Error
	}

	return trx, nil
}

func (r *repotrx) Delete(id uint) error {

	result := r.DB.Delete(&model.Trx{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
