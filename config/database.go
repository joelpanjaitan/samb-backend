package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DataBase*gorm.DB

func InitDB() *gorm.DB {
	network := "root:password@tcp(localhost:3307)/db_company?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(network), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection failed to database: %v", err)
	}
	
	DataBase= DB
	return DB
}