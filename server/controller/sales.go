package controller

import (
	"net/http"

	"github.com/aman-lf/sales-server/service"
	"github.com/gin-gonic/gin"
)

func SetupsaleRoute(router *gin.Engine) {
	router.GET("/api/sale", GetsalesHandler)
	router.GET("/api/sale-product", GetSalesByProductHandler)
	router.GET("/api/sale-brand", GetSalesByBrandHandler)
}

func GetsalesHandler(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	sales, err := service.GetSales(c, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    sales,
		"message": "Successfully retrieved sales",
	})
}

func GetSalesByProductHandler(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	searchText := c.Query("searchText")

	sales, err := service.GetSalesByProduct(c, limit, offset, searchText)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    sales,
		"message": "Successfully retrieved sales",
	})
}

func GetSalesByBrandHandler(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	searchText := c.Query("searchText")

	sales, err := service.GetSalesByBrand(c, limit, offset, searchText)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    sales,
		"message": "Successfully retrieved sales",
	})
}
