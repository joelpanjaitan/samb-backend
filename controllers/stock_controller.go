package controllers

import (
	"net/http"

	"samb-backend/config"

	"github.com/gin-gonic/gin"
)

func GetStockReport(c *gin.Context) {
	var stock []struct {
		Warehouse string
		Product   string
		QtyDus    int
		QtyPcs    int
	}

	query := `
		SELECT 
			w.name as Warehouse, 
			p.name as Product, 
			COALESCE(SUM(d.qty_dus), 0) as QtyDus, 
			COALESCE(SUM(d.qty_pcs), 0) as QtyPcs
		FROM warehouses w
		JOIN products p
		LEFT JOIN trx_in_details d ON p.id = d.product_id
		GROUP BY w.name, p.name
	`

	config.DataBase.Raw(query).Scan(&stock)
	c.JSON(http.StatusOK, stock)
}