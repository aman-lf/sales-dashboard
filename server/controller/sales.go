package controller

import (
	"net/http"

	"github.com/aman-lf/sales-server/service"
	"github.com/aman-lf/sales-server/utils"
	"github.com/gin-gonic/gin"
)

func SetupSaleRoute(router *gin.Engine) {
	router.GET("/api/sale", GetsalesHandler)
	router.GET("/api/sale-product", GetSalesByProductHandler)
	router.GET("/api/sale-brand", GetSalesByBrandHandler)
	router.GET("/api/dashboard", GetDashboardSalesHandler)
}

// @Summary Get all sales
// @Description Retrieve all sales data
// @Produce json
// @Success 200 {object} object "Successfully retrieved sales"
// @Router /api/sale [get]
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

// @Summary Get sales by product
// @Description Retrieve sales data grouped by product
// @Produce json
// @Param perPage query int false "Items per page"
// @Param page query int false "Page number"
// @Param sortBy query string false "Field to sort by"
// @Param sortOrder query int false "Sort order (asc/desc)"
// @Param searchText query string false "Text to search"
// @Success 200 {object} object "Successfully retrieved sales by product"
// @Router /api/sale-product [get]
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
		"message": "Successfully retrieved sales by product",
	})
}

// @Summary Get sales by brand
// @Description Retrieve sales data grouped by brand
// @Produce json
// @Param perPage query int false "Items per page"
// @Param page query int false "Page number"
// @Param sortBy query string false "Field to sort by"
// @Param sortOrder query int false "Sort order (asc/desc)"
// @Param searchText query string false "Text to search"
// @Success 200 {object} object "Successfully retrieved sales by brand"
// @Router /api/sale-brand [get]
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
		"message": "Successfully retrieved sales by brand",
	})
}

// @Summary Get dashboard sales data
// @Description Retrieve sales data for the dashboard
// @Produce json
// @Success 200 {object} object "Successfully retrieved sales dashboard"
// @Router /api/dashboard [get]
func GetDashboardSalesHandler(c *gin.Context) {
	data, err := service.GetDashboardData(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "Successfully retrieved sales dashboard",
	})
}
