package repo

import (
	"errors"
	"shop/model"

	"gorm.io/gorm"
)

// TrxID        uint
// LogProductID uint
// TokoID       uint

type FotoProduct interface {
	Create(fotoproduct model.FotoProduct) (model.FotoProduct, error)
	GetByID(fotoproductid uint) (model.FotoProduct, error)
	GetByProductID(productid uint) ([]model.FotoProduct, error)

	Update(fotoproduct model.FotoProduct) (model.FotoProduct, error)
	Delete(fotoproductid uint) error
}

func NewFotoProductRepo(db *gorm.DB) FotoProduct {
	return &repofotoproduct{
		DB: db,
	}
}

type repofotoproduct struct {
	DB *gorm.DB
}

func (r *repofotoproduct) Create(fotoproduct model.FotoProduct) (model.FotoProduct, error) {

	result := r.DB.Create(&fotoproduct)
	if result.Error != nil {
		return model.FotoProduct{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.FotoProduct{}, errors.New("not created")
	}

	return fotoproduct, nil
}

func (r *repofotoproduct) GetByID(id uint) (model.FotoProduct, error) {
	var fotoproduct model.FotoProduct
	result := r.DB.First(&fotoproduct, id)
	if result.Error != nil {
		return model.FotoProduct{}, result.Error
	}

	return fotoproduct, nil
}

func (r *repofotoproduct) GetByProductID(productid uint) ([]model.FotoProduct, error) {
	var fotoproducts []model.FotoProduct
	result := r.DB.Where(&model.FotoProduct{ProductID: productid}).Find(&fotoproducts)
	if result.Error != nil {
		return []model.FotoProduct{}, result.Error
	}

	return fotoproducts, nil
}

func (r *repofotoproduct) Update(fotoproduct model.FotoProduct) (model.FotoProduct, error) {

	var olddata model.FotoProduct

	result := r.DB.First(&olddata, fotoproduct.ID)
	if result.Error != nil {
		return model.FotoProduct{}, result.Error
	}

	fotoproduct.CreatedAt = olddata.CreatedAt

	result = r.DB.Create(&fotoproduct)
	if result.Error != nil {
		return model.FotoProduct{}, result.Error
	}

	return fotoproduct, nil
}

func (r *repofotoproduct) Delete(id uint) error {

	result := r.DB.Delete(&model.FotoProduct{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
