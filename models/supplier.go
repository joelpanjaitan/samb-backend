package models

type Supplier struct {
	SupplierPK   int    `gorm:"primaryKey"`
	SupplierName string `gorm:"size:255;not null"`
}