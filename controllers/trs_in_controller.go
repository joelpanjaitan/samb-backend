package controllers

import (
	"fmt"
	"net/http"
	"samb-backend/config"
	"samb-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateTrxIn(c *gin.Context) {
	var header models.TrxInHeader
	if err := c.ShouldBindJSON(&header); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	result := config.DataBase.Exec(
		`INSERT INTO TransaksiPenerimaanBarangHeader (TrxInNo, WhsIdf, TrxInDate, TrxInSuppIdf, TrxInNotes) 
		VALUES (?, ?, ?, ?, ?)`,
		header.TrxInNo, header.WhsIdf, header.TrxInDate, header.TrxInSuppIdf, header.TrxInNotes,
	)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	lastCreatedID := header.TrxInPK
	for _, detail := range header.Details {
		detail.TrxInIDF = lastCreatedID 
		err := config.DataBase.Exec(
			`INSERT INTO TransaksiPenerimaanBarangDetail (TrxInIDF, TrxInDProductIdf, TrxInDQtyDus, TrxInDQtyPcs) 
			VALUES (?, ?, ?, ?)`,
			detail.TrxInIDF, detail.TrxInDProductIdf, detail.TrxInDQtyDus, detail.TrxInDQtyPcs,
		).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction Created"})
}

func GetTrxIn(c *gin.Context) {
	var headers []models.TrxInHeader
	rows,err := config.DataBase.Raw(`SELECT * FROM TransaksiPenerimaanBarangHeader`,).Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error fetching products data: %s", err.Error()),
		})
		return 
	}
	defer rows.Close()
	for rows.Next() {
		var header models.TrxInHeader
		if err := rows.Scan(&header.TrxInPK, &header.TrxInNo, &header.WhsIdf, &header.TrxInDate, &header.TrxInSuppIdf, &header.TrxInNotes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Fetch details for each header
		var details []models.TrxInDetail
		detailRows, err := config.DataBase.Raw(
			`SELECT * FROM TransaksiPenerimaanBarangDetail WHERE TrxInIDF = ?`,
			header.TrxInPK,
		).Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer detailRows.Close()

		for detailRows.Next() {
			var detail models.TrxInDetail
			if err := detailRows.Scan(&detail.TrxInDPK, &detail.TrxInIDF, &detail.TrxInDProductIdf, &detail.TrxInDQtyDus, &detail.TrxInDQtyPcs); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			details = append(details, detail)
		}

		header.Details = details
		headers = append(headers, header)
	}

	c.JSON(http.StatusOK, headers)
}
