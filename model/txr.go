package model

type TrxComplete struct {
	Trx       Trx
	DetailTrx DetailTrx
}

type Trx struct {
	BasicModelShowID
	HargaTotal  int    `gorm:"column:harga_total;type:int" json:"harga_total"`
	KodeInvoice string `gorm:"column:kode_invoice;type:varchar(255)" json:"kode_invoice"`
	MethodBayar string `gorm:"column:method_bayar;type:varchar(255)" json:"method_bayar"`

	UserID uint

	AlamatID uint
	Alamat   Alamat
}
