package controller

import (
	"net/http"

	"github.com/aman-lf/sales-server/service"
	"github.com/gin-gonic/gin"
)

func SetupsaleRoute(router *gin.Engine) {
	r := router.Group("/api/sale")
	r.GET("/", GetsalesHandler)
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
		"sales":   sales,
		"message": "Successfully retrieved sales",
	})
}
