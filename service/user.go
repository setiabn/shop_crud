package service

import (
	"shop/model"
	"shop/repo"
)

type User interface {
	Get(userid uint) (model.User, error)
	Update(user model.User) (model.User, error)
}

func NewServiceUser(repo repo.User) User {
	return &serviceUser{
		repo: repo,
	}
}

type serviceUser struct {
	repo repo.User
}

func (s *serviceUser) Get(userid uint) (model.User, error) {

	newUser, err := s.repo.Get(userid)
	if err != nil {
		return model.User{}, err
	}
	return newUser, nil
}

func (s *serviceUser) Update(user model.User) (model.User, error) {
	newUser, err := s.repo.Update(user)
	if err != nil {
		return model.User{}, err
	}
	return newUser, nil
}
