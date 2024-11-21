package controllers

import (
	"net/http"

	"samb-backend/config"
	"samb-backend/models"

	"github.com/gin-gonic/gin"
)

func GetSuppliers(c *gin.Context) {
	var suppliers []models.Supplier
	config.DataBase.Find(&suppliers)
	c.JSON(http.StatusOK, suppliers)
}

func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	config.DataBase.Find(&customers)
	c.JSON(http.StatusOK, customers)
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DataBase.Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetWarehouses(c *gin.Context) {
	var warehouses []models.Warehouse
	config.DataBase.Find(&warehouses)
	c.JSON(http.StatusOK, warehouses)
}