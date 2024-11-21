package controllers

import (
	"net/http"

	"samb-backend/config"

	"github.com/gin-gonic/gin"
)

func GetStockReport(c *gin.Context) {
	var stockReport []struct {
		WarehouseName string `json:"warehouse"`
		ProductName   string `json:"product"`
		QtyDus        int    `json:"qty_dus"`
		QtyPcs        int    `json:"qty_pcs"`
	}

	query := `
		SELECT 
			w.whs_name AS warehouse, 
			p.product_name AS product, 
			COALESCE(SUM(in_d.trx_in_d_qty_dus), 0) - COALESCE(SUM(out_d.trx_out_d_qty_dus), 0) AS qty_dus,
			COALESCE(SUM(in_d.trx_in_d_qty_pcs), 0) - COALESCE(SUM(out_d.trx_out_d_qty_pcs), 0) AS qty_pcs
		FROM 
			master_warehouse w
		LEFT JOIN master_product p ON TRUE
		LEFT JOIN trx_in_detail in_d ON in_d.trx_in_d_product_idf = p.product_pk
		LEFT JOIN trx_out_detail out_d ON out_d.trx_out_d_product_idf = p.product_pk
		GROUP BY w.whs_name, p.product_name
	`

	config.DataBase.Raw(query).Scan(&stockReport)
	c.JSON(http.StatusOK, stockReport)
}