package models

type TrxInHeader struct {
	ID      int    `gorm:"primaryKey"`
	Number  string `gorm:"size:255;not null"`
	WhsID   int
	Date    string `gorm:"not null"`
	SuppID  int
	Notes   string
	Details []TrxInDetail `gorm:"foreignKey:HeaderID"`
}

type TrxInDetail struct {
	ID        int `gorm:"primaryKey"`
	HeaderID  int
	ProductID int
	QtyDus    int
	QtyPcs    int
}
