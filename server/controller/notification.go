package controller

import (
	"github.com/aman-lf/sales-server/service"
	"github.com/gin-gonic/gin"
)

func SetupNotificationRoute(router *gin.Engine) {
	router.GET("/events", HandleSSEConnection)
}

func HandleSSEConnection(ctx *gin.Context) {
	// Set response headers for SSE
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	c := service.NewSSEController()
	// Add client to the list of SSE clients
	c.AddClient(ctx)
	defer c.RemoveClient(ctx)

	// Wait indefinitely (or until the client disconnects)
	<-ctx.Done()
}
