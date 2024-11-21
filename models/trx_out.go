package models

type TrxOutHeader struct {
	ID      int    `gorm:"primaryKey"`
	Number  string `gorm:"size:255;not null"`
	WhsID   int
	Date    string `gorm:"not null"`
	SuppID  int
	Notes   string
	Details []TrxOutDetail `gorm:"foreignKey:HeaderID"`
}

type TrxOutDetail struct {
	ID        int `gorm:"primaryKey"`
	HeaderID  int
	ProductID int
	QtyDus    int
	QtyPcs    int
}
