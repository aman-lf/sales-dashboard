package controller

import (
	"github.com/aman-lf/sales-server/service"
	"github.com/gin-gonic/gin"
)

func SetupNotificationRoute(router *gin.Engine) {
	router.GET("/new-file-notification", SSEConnectionHandler)
}

// SSEConnectionHandler godoc
// @Summary Establish a Server-Sent Events (SSE) connection
// @Description Initiates a connection for Server-Sent Events (SSE) with the client
// @Produce text/plain
// @Router /new-file-notification [get]
func SSEConnectionHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Access-Control-Allow-Origin", "*")

	service.SetupSSEMessage(ctx)
}
