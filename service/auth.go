package service

import (
	"shop/model"
	"shop/repo"

	"github.com/gofiber/fiber/v2"
)

type Auth interface {
	Login(user model.User) (model.User, error)
	Register(user model.User) error
}

func NewServiceAuth(repoUser repo.User, repoToko repo.Toko, repoTrx repo.Trx, repoAlamat repo.Alamat, repoProvCity repo.ProvCity) Auth {
	return &serviceAuth{
		rUser: repoUser,
		rToko: repoToko,
	}
}

type serviceAuth struct {
	rUser repo.User
	rToko repo.Toko
}

func (s *serviceAuth) Register(user model.User) error {

	var err error

	hashPass, err := hashPassword(user.KataSandi)
	if err != nil {
		return err
	}
	user.KataSandi = hashPass

	user.Toko = model.Toko{
		NamaToko: "toko " + user.Nama,
		UserID:   user.ID,
	}

	_, err = s.rUser.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceAuth) Login(user model.User) (model.User, error) {

	savedUser, err := s.rUser.GetByNoTelp(user.NoTelp)
	if err != nil {
		return model.User{}, err
	}

	toko, err := s.rToko.GetByUserID(savedUser.ID)
	if err != nil {
		return model.User{}, err
	}

	savedUser.Toko = toko
	// savedUser.Trxs = trxs
	// savedUser.Alamats = alamats

	if !CheckPasswordHash(user.KataSandi, savedUser.KataSandi) {
		return model.User{}, fiber.ErrUnauthorized
	}

	return savedUser, nil
}
