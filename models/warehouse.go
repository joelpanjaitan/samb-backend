package models

type Warehouse struct {
	WhsPK   int    `gorm:"primaryKey"`
	WhsName string `gorm:"size:255;not null"`
}