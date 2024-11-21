package controllers

import (
	"net/http"

	"samb-backend/config"
	"samb-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateTrxIn(c *gin.Context) {
	var header models.TrxInHeader
	if err := c.ShouldBindJSON(&header); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DataBase.Create(&header)
	c.JSON(http.StatusOK, gin.H{"message": "Transaction Created"})
}

func GetTrxIn(c *gin.Context) {
	var headers []models.TrxInHeader
	config.DataBase.Preload("Details").Find(&headers)
	c.JSON(http.StatusOK, headers)
}
