package models

type Product struct {
	ProductPK   int    `gorm:"primaryKey"`
	ProductName string `gorm:"size:255;not null"`
}