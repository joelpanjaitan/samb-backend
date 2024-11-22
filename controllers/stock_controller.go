package controllers

import (
	"net/http"

	"samb-backend/config"
	"samb-backend/models"

	"github.com/gin-gonic/gin"
)

func GetStockReport(c *gin.Context) {
	var stockReport []models.Stock

	query := `
		SELECT 
			w.WhsName AS warehouse, 
			p.ProductName AS product, 
			COALESCE(SUM(in_d.TrxInDQtyDus), 0) - COALESCE(SUM(out_d.TrxOutDQtyDus), 0) AS qty_dus,
			COALESCE(SUM(in_d.TrxInDQtyPcs), 0) - COALESCE(SUM(out_d.TrxOutDQtyPcs), 0) AS qty_pcs
		FROM 
			MasterWarehouse w
		LEFT JOIN TransaksiPenerimaanBarangHeader in_h ON in_h.WhsIdf = w.WhsPK
		LEFT JOIN TransaksiPenerimaanBarangDetail in_d ON in_d.TrxInIDF = in_h.TrxInPK
		LEFT JOIN TransaksiPengeluaranBarangHeader out_h ON out_h.WhsIdf = w.WhsPK
		LEFT JOIN TransaksiPengeluaranBarangDetail out_d ON out_d.TrxOutIDF = out_h.TrxOutPK
		LEFT JOIN MasterProduct p ON p.ProductPK = in_d.TrxInDProductIdf OR p.ProductPK = out_d.TrxOutDProductIdf
		GROUP BY w.WhsName, p.ProductName
	`

	if err := config.DataBase.Raw(query).Scan(&stockReport).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Database error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stockReport)
}