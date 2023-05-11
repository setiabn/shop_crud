package repo

import (
	"errors"
	"shop/model"

	"gorm.io/gorm"
)

type User interface {
	Create(user model.User) (model.User, error)
	Get(userid uint) (model.User, error)
	GetByNoTelp(notelp string) (model.User, error)

	Update(user model.User) (model.User, error)
	Delete(userid uint) error
}

func NewUserRepo(db *gorm.DB) User {
	return &repouser{
		DB: db,
	}
}

type repouser struct {
	DB *gorm.DB
}

func (r *repouser) Create(user model.User) (model.User, error) {

	result := r.DB.Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.User{}, errors.New("not created")
	}

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (r *repouser) Get(id uint) (model.User, error) {
	var user model.User

	result := r.DB.First(&user, id)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (r *repouser) GetByNoTelp(notelp string) (model.User, error) {
	var user model.User

	result := r.DB.Where(&model.User{NoTelp: notelp}).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (r *repouser) Update(user model.User) (model.User, error) {

	var olddata model.User

	result := r.DB.First(&olddata, user.ID)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	user.CreatedAt = olddata.CreatedAt
	user.IsAdmin = olddata.IsAdmin
	user.JenisKelamin = olddata.JenisKelamin
	user.Tentang = olddata.Tentang

	result = r.DB.Save(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (r *repouser) Delete(id uint) error {

	result := r.DB.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
