package models

type Warehouse struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:255;not null"`
}