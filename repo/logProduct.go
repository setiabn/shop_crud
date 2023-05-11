package repo

import (
	"errors"
	"shop/model"

	"gorm.io/gorm"
)

type LogProduct interface {
	Create(logproduct model.LogProduct) (model.LogProduct, error)
	GetByID(logproductid uint) (model.LogProduct, error)
	GetByProductID(productid uint) (model.LogProduct, error)

	Update(logproduct model.LogProduct) (model.LogProduct, error)
	Delete(logproductid uint) error
}

func NewLogProductRepo(db *gorm.DB) LogProduct {
	return &repologproduct{
		DB: db,
	}
}

type repologproduct struct {
	DB *gorm.DB
}

func (r *repologproduct) Create(logproduct model.LogProduct) (model.LogProduct, error) {

	result := r.DB.Create(&logproduct)
	if result.Error != nil {
		return model.LogProduct{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.LogProduct{}, errors.New("not created")
	}

	return logproduct, nil
}

func (r *repologproduct) GetByID(id uint) (model.LogProduct, error) {
	var logproduct model.LogProduct
	result := r.DB.First(&logproduct, id)
	if result.Error != nil {
		return model.LogProduct{}, result.Error
	}

	return logproduct, nil
}

func (r *repologproduct) GetByProductID(productid uint) (model.LogProduct, error) {
	var logproduct model.LogProduct
	result := r.DB.Where(&model.LogProduct{ProductID: productid}).First(&logproduct)
	if result.Error != nil {
		return model.LogProduct{}, result.Error
	}

	return logproduct, nil
}

func (r *repologproduct) Update(logproduct model.LogProduct) (model.LogProduct, error) {

	var olddata model.LogProduct

	result := r.DB.First(&olddata, logproduct.ID)
	if result.Error != nil {
		return model.LogProduct{}, result.Error
	}

	logproduct.CreatedAt = olddata.CreatedAt

	result = r.DB.Create(&logproduct)
	if result.Error != nil {
		return model.LogProduct{}, result.Error
	}

	return logproduct, nil
}

func (r *repologproduct) Delete(id uint) error {

	result := r.DB.Delete(&model.LogProduct{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
