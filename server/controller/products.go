package controller

import (
	"net/http"

	"github.com/aman-lf/sales-server/service"
	"github.com/gin-gonic/gin"
)

func SetupProductRoute(router *gin.Engine) {
	r := router.Group("/api/product")
	r.GET("/", GetProductsHandler)
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
		"products": products,
		"message":  "Successfully retrieved products",
	})
}
