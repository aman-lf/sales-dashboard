package controller

import (
	"net/http"

	"github.com/aman-lf/sales-server/service"
	"github.com/gin-gonic/gin"
)

func SetupProductRoute(router *gin.Engine) {
	router.GET("/api/product", GetProductsHandler)
}

func GetProductsHandler(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	products, err := service.GetProducts(c, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    products,
		"message": "Successfully retrieved products",
	})
}
