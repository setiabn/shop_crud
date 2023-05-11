package service

import (
	"shop/model"
	"shop/repo"
)

type User interface {
	Get(userid uint) (model.User, error)
	Update(user model.User) (model.User, error)
}

func NewServiceUser(repoUser repo.User, repoToko repo.Toko, repoTrx repo.Trx, repoAlamat repo.Alamat, repoProvCity repo.ProvCity) User {
	return &serviceUser{
		repo:      repoUser,
		rToko:     repoToko,
		rTrx:      repoTrx,
		rAlamat:   repoAlamat,
		rProvCity: repoProvCity,
	}
}

type serviceUser struct {
	repo      repo.User
	rToko     repo.Toko
	rTrx      repo.Trx
	rAlamat   repo.Alamat
	rProvCity repo.ProvCity
}

func (s *serviceUser) Get(userid uint) (model.User, error) {

	user, err := s.repo.Get(userid)
	if err != nil {
		return model.User{}, err
	}

	toko, err := s.rToko.GetByUserID(user.ID)
	if err != nil {
		return model.User{}, err
	}

	trxs, err := s.rTrx.GetByUserID(user.ID)
	if err != nil {
		if err.Error() != "record not found" {
			return model.User{}, err
		}
	}

	alamats, err := s.rAlamat.GetByUserID(user.ID)
	if err != nil {
		if err.Error() != "record not found" {
			return model.User{}, err
		}
	}

	user.Toko = toko
	user.Trxs = trxs
	user.Alamats = alamats

	return user, nil
}

func (s *serviceUser) Update(user model.User) (model.User, error) {
	newUser, err := s.repo.Update(user)
	if err != nil {
		return model.User{}, err
	}
	return newUser, nil
}
