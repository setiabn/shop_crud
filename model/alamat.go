package model

type Alamat struct {
	BasicModelShowID

	JudulAlamat  string `gorm:"column:judul_alamat;type:varchar(255)" json:"judul_alamat" validate:"required"`
	NamaPenerima string `gorm:"column:nama_penerima;type:varchar(255)" json:"nama_penerima" validate:"required"`
	NoTelp       string `gorm:"column:no_telp;type:varchar(255)" json:"no_telp" validate:"required"`
	DetailAlamat string `gorm:"column:detail_alamat;type:varchar(255)" json:"detail_alamat" validate:"required"`

	UserID uint `gorm:"column:user_id;type:int" json:"-"`
}
