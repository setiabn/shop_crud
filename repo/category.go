package repo

import (
	"errors"
	"shop/model"

	"gorm.io/gorm"
)

type Category interface {
	Create(category model.Category) (model.Category, error)
	Get(categoryid uint) (model.Category, error)
	GetAll() ([]model.Category, error)

	Update(category model.Category) (model.Category, error)
	Delete(categoryid uint) error
}

func NewCategoryRepo(db *gorm.DB) Category {
	return &repocategory{
		DB: db,
	}
}

type repocategory struct {
	DB *gorm.DB
}

func (r *repocategory) Create(category model.Category) (model.Category, error) {

	result := r.DB.Create(&category)
	if result.Error != nil {
		return model.Category{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.Category{}, errors.New("not created")
	}

	return category, nil
}

func (r *repocategory) Get(id uint) (model.Category, error) {
	var category model.Category
	result := r.DB.First(&category, id)
	if result.Error != nil {
		return model.Category{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.Category{}, errRecordNotFound
	}

	return category, nil
}

func (r *repocategory) GetAll() ([]model.Category, error) {

	var categories []model.Category
	result := r.DB.Find(&categories)
	if result.Error != nil {
		return []model.Category{}, result.Error
	}
	if result.RowsAffected == 0 {
		return []model.Category{}, errRecordNotFound
	}

	return categories, nil
}

func (r *repocategory) Update(category model.Category) (model.Category, error) {

	var olddata model.Category

	result := r.DB.First(&olddata, category.ID)
	if result.Error != nil {
		return model.Category{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.Category{}, errRecordNotFound
	}

	category.CreatedAt = olddata.CreatedAt

	result = r.DB.Save(&category)
	if result.Error != nil {
		return model.Category{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.Category{}, errRecordNotFound
	}

	return category, nil
}

func (r *repocategory) Delete(id uint) error {

	result := r.DB.Delete(&model.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errRecordNotFound
	}

	return nil
}
