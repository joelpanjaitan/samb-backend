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
