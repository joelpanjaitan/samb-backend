package routes

import (
	"samb-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// For master data
	router.GET("/suppliers", controllers.GetSuppliers)
	router.GET("/customers", controllers.GetCustomers)
	router.GET("/products", controllers.GetProducts)
	router.GET("/warehouses", controllers.GetWarehouses)

	// For transactions
	router.POST("/trx-in", controllers.CreateTrxIn)
	router.GET("/trx-in", controllers.GetTrxIn)

	router.POST("/trx-out", controllers.CreateTrxOut)
	router.GET("/trx-out", controllers.GetTrxOut)

	// For stock report
	router.GET("/stock", controllers.GetStockReport)
}
