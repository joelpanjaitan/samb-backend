package models

type Customer struct {
	CustomerPK   int    `gorm:"primaryKey"`
	CustomerName string `gorm:"size:255;not null"`
}