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
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	
	result := config.DataBase.Exec(
		`INSERT INTO TransaksiPengeluaranBarangHeader (TrxOutNo, WhsIdf, TrxOutDate, TrxOutSuppIdf, TrxOutNotes) 
		VALUES (?, ?, ?, ?, ?)`,
		header.TrxOutNo, header.WhsIdf, header.TrxOutDate, header.TrxOutSuppIdf, header.TrxOutNotes,
	)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		return
	}

	lastCreatedID := header.TrxOutPK

	for _, detail := range header.Details {
		detail.TrxOutIDF = lastCreatedID // Link detail to the header
		err := config.DataBase.Exec(
			`INSERT INTO TransaksiPengeluaranBarangDetail (TrxOutIDF, TrxOutDProductIdf, TrxOutDQtyDus, TrxOutDQtyPcs) 
			VALUES (?, ?, ?, ?)`,
			detail.TrxOutIDF, detail.TrxOutDProductIdf, detail.TrxOutDQtyDus, detail.TrxOutDQtyPcs,
		).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction succesfully created"})
}

func GetTrxOut (c*gin.Context){
	var transactions []models.TrxOutHeader
	rows, err := config.DataBase.Raw(
		`SELECT * FROM TransaksiPengeluaranBarangHeader`,
	).Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var header models.TrxOutHeader
		if err := rows.Scan(&header.TrxOutPK, &header.TrxOutNo, &header.WhsIdf, &header.TrxOutDate, &header.TrxOutSuppIdf, &header.TrxOutNotes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var details []models.TrxOutDetail
		detailRows, err := config.DataBase.Raw(
			`SELECT * FROM TransaksiPengeluaranBarangDetail WHERE TrxOutIDF = ?`,
			header.TrxOutPK,
		).Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer detailRows.Close()

		for detailRows.Next() {
			var detail models.TrxOutDetail
			if err := detailRows.Scan(&detail.TrxOutDPK, &detail.TrxOutIDF, &detail.TrxOutDProductIdf, &detail.TrxOutDQtyDus, &detail.TrxOutDQtyPcs); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			details = append(details, detail)
		}

		header.Details = details
		transactions = append(transactions, header)
	}
	c.JSON(http.StatusOK, transactions)
}