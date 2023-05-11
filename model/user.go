package model

import (
	"time"
)

type User struct {
	BasicModelShowID

	Nama string `gorm:"column:nama;type:varchar(255)" json:"nama_category" validate:"required"`

	// hide from json
	KataSandi string `gorm:"column:kata_sandi;type:varchar(255)" json:"-" validate:"required"`

	NoTelp       string    `gorm:"column:no_telp;type:varchar(255);unique" json:"no_telp" validate:"required"`
	TanggalLahir time.Time `gorm:"column:tanggal_lahir;type:datetime" json:"tanggal_lahir" validate:"required"`
	JenisKelamin string    `gorm:"column:jenis_kelamin;type:varchar(255)" json:"jenis_kelamin" validate:"required"`
	Tentang      string    `gorm:"column:tentang;type:varchar(255)" json:"tentang" validate:"required"`
	Pekerjaan    string    `gorm:"column:pekerjaan;type:varchar(255)" json:"pekerjaan" validate:"required"`
	Email        string    `gorm:"column:email;type:varchar(255)" json:"email" validate:"required"`
	IDProvinsi   string    `gorm:"column:id_provinsi;type:varchar(255)" json:"id_provinsi" validate:"required"`
	IDKota       string    `gorm:"column:id_kota;type:varchar(255)" json:"id_kota" validate:"required"`
	IsAdmin      bool      `gorm:"column:is_admin" json:"-"`

	Toko    Toko     `json:"-"`
	Alamats []Alamat `json:"-"`
	Trxs    []Trx    `json:"-"`
}
