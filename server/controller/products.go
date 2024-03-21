package controller

import (
	"net/http"

	"github.com/aman-lf/sales-server/service"
	"github.com/aman-lf/sales-server/utils/logger"
	"github.com/gin-gonic/gin"
)

func SetupProductRoute(router *gin.Engine) {
	router.GET("/api/product", GetProductsHandler)
}

// GetProductsHandler godoc
// @Summary Get all products
// @Description Retrieve all products data
// @Produce json
// @Param limit query int false "Limit number of products per page"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} object "Successfully retrieved products"
// @Failure 500 {object} object "Internal Server Error"
// @Router /api/product [get]
func GetProductsHandler(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	products, err := service.GetProducts(c, limit, offset)
	if err != nil {
		logger.Error("Failed to get products", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    products,
		"message": "Successfully retrieved products",
	})
}
