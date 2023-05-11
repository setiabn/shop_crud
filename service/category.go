package service

import (
	"shop/model"
	"shop/repo"
)

type Category interface {
	Create(category model.Category) (model.Category, error)
	Get(categoryid uint) (model.Category, error)
	GetAll() ([]model.Category, error)
	Update(category model.Category) (model.Category, error)
	Delete(categoryid uint) error
}

func NewServiceCategory(repo repo.Category) Category {
	return &serviceCategory{
		repo: repo,
	}
}

type serviceCategory struct {
	repo repo.Category
}

func (s *serviceCategory) Create(category model.Category) (model.Category, error) {

	newCategory, err := s.repo.Create(category)
	if err != nil {
		return model.Category{}, err
	}
	return newCategory, nil
}

func (s *serviceCategory) Get(categoryid uint) (model.Category, error) {
	newCategory, err := s.repo.Get(categoryid)
	if err != nil {
		return model.Category{}, err
	}
	return newCategory, nil
}

func (s *serviceCategory) GetAll() ([]model.Category, error) {
	allCategory, err := s.repo.GetAll()
	if err != nil {
		return []model.Category{}, err
	}
	return allCategory, nil
}

func (s *serviceCategory) Update(category model.Category) (model.Category, error) {
	newCategory, err := s.repo.Update(category)
	if err != nil {
		return model.Category{}, err
	}
	return newCategory, nil
}

func (s *serviceCategory) Delete(categoryid uint) error {
	err := s.repo.Delete(categoryid)
	if err != nil {
		return err
	}
	return nil
}
