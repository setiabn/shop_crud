package repo

import (
	"errors"
	"shop/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Product interface {
	Create(product model.Product) (model.Product, error)
	Get(productid uint) (model.Product, error)
	GetAll(limit uint, page uint, categoryid uint, tokoid uint) ([]model.Product, error)
	Update(product model.Product) (model.Product, error)
	Delete(productid uint) error
}

func NewProductRepo(db *gorm.DB) Product {
	return &repoproduct{
		DB: db,
	}
}

type repoproduct struct {
	DB *gorm.DB
}

func (r *repoproduct) Create(product model.Product) (model.Product, error) {

	result := r.DB.Create(&product)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.Product{}, errors.New("not created")
	}

	return product, nil
}

func (r *repoproduct) Get(id uint) (model.Product, error) {
	var product model.Product
	result := r.DB.First(&product, id)
	if result.Error != nil {
		return model.Product{}, result.Error
	}

	return product, nil
}

func (r *repoproduct) GetAll(limit uint, page uint, categoryid uint, tokoid uint) ([]model.Product, error) {
	var products []model.Product

	tx := r.DB

	if limit > 0 {
		tx = tx.Limit(int(limit))
	}

	if page > 0 {
		offset := int(limit) * int(page-1)
		tx = tx.Offset(offset)
	}

	if categoryid > 0 {
		tx = tx.Where(&model.Product{CategoryID: categoryid})
	}

	if tokoid > 0 {
		tx = tx.Where(&model.Product{TokoID: tokoid})
	}

	result := tx.Find(&products)

	if result.Error != nil {
		return []model.Product{}, result.Error
	}
	return products, nil
}

func (r *repoproduct) Update(product model.Product) (model.Product, error) {

	var olddata model.Product

	result := r.DB.First(&olddata, product.ID)
	if result.Error != nil {
		return model.Product{}, result.Error
	}

	product.CreatedAt = olddata.CreatedAt
	if olddata.TokoID != product.TokoID || olddata.Toko.UserID != product.Toko.UserID {
		return model.Product{}, fiber.ErrUnauthorized
	}

	result = r.DB.Save(&product)
	if result.Error != nil {
		return model.Product{}, result.Error
	}

	return product, nil
}

func (r *repoproduct) Delete(id uint) error {

	var oldproduct model.Product
	if err := r.DB.First(&oldproduct, id).Error; err != nil {
		return err
	}

	result := r.DB.Delete(&model.Product{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
