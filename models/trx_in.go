package models

type TrxInHeader struct {
	TrxInPK      int    `gorm:"primaryKey"`
	TrxInNo      string `gorm:"size:255;not null"`
	WhsIdf       int
	TrxInDate    string `gorm:"not null"`
	TrxInSuppIdf int
	TrxInNotes   string
	Details      []TrxInDetail `gorm:"foreignKey:HeaderID"`
}

type TrxInDetail struct {
	TrxInDPK         int `gorm:"primaryKey"`
	TrxInIDF         int
	TrxInDProductIdf int
	TrxInDQtyDus     int
	TrxInDQtyPcs     int
}
