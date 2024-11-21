package models

type Supplier struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:255;not null"`
}