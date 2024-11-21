package models

type TrxOutHeader struct {
	TrxOutPK      int    `gorm:"primaryKey"`
	TrxOutNo      string `gorm:"size:255;not null"`
	WhsIdf        int
	TrxOutDate    string `gorm:"not null"`
	TrxOutSuppIdf int
	TrxOutNotes   string
	Details       []TrxOutDetail `gorm:"foreignKey:HeaderID"`
}

type TrxOutDetail struct {
	TrxOutDPK         int `gorm:"primaryKey"`
	TrxOutIDF         int
	TrxOutDProductIdf int
	TrxOutDQtyDus     int
	TrxOutDQtyPcs     int
}
