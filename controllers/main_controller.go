package controllers

import (
	"fmt"
	"net/http"
	"samb-backend/config"
	"samb-backend/models"

	"github.com/gin-gonic/gin"
)

func GetSuppliers(c *gin.Context) {
	var suppliers []models.Supplier
	err := config.DataBase.Raw("SELECT * FROM MasterSupplier").Scan(&suppliers).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error fetching suppliers data: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, suppliers)
}

func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	err := config.DataBase.Raw("SELECT * FROM MasterCustomer").Scan(&customers).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error fetching customers data: %s", err.Error()),
		})
		return 
	}
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