package controllers

import (
	"net/http"
	"samb-backend/config"
	"samb-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateTrxOut(c*gin.Context){
	var header models.TrxOutHeader
	if err := c.ShouldBindJSON(&header); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DataBase.Create(&header).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, header)
}

func GetTrxOut (c*gin.Context){
	var transactions []models.TrxOutHeader
	config.DataBase.Preload("Details").Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}