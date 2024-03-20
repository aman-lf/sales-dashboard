package controller

import (
	"net/http"

	"github.com/aman-lf/sales-server/service"
	"github.com/aman-lf/sales-server/utils"
	"github.com/gin-gonic/gin"
)

func SetupsaleRoute(router *gin.Engine) {
	router.GET("/api/sale", GetsalesHandler)
	router.GET("/api/sale-product", GetSalesByProductHandler)
	router.GET("/api/sale-brand", GetSalesByBrandHandler)
	router.GET("/api/dashboard", GetDashboardSalesHandler)
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
	limit := c.Query("perPage")
	page := c.Query("page")
	sortBy := c.Query("sortBy")
	sortOrder := c.Query("sortOrder")
	searchText := c.Query("searchText")
	pipelineFilter := utils.GetPipelineFilter(limit, page, sortBy, service.PRODUCT_DEFAULT_SORT, sortOrder, searchText)

	data, err := service.GetSalesByProduct(c, pipelineFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"meta":    pipelineFilter,
		"message": "Successfully retrieved sales",
	})
}

func GetSalesByBrandHandler(c *gin.Context) {
	limit := c.Query("perPage")
	page := c.Query("page")
	sortBy := c.Query("sortBy")
	sortOrder := c.Query("sortOrder")
	searchText := c.Query("searchText")
	pipelineFilter := utils.GetPipelineFilter(limit, page, sortBy, service.BRAND_DEFAULT_SORT, sortOrder, searchText)

	data, err := service.GetSalesByBrand(c, pipelineFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"meta":    pipelineFilter,
		"message": "Successfully retrieved sales",
	})
}

func GetDashboardSalesHandler(c *gin.Context) {
	data, err := service.GetDashboardData(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "Successfully retrieved sales",
	})
}
