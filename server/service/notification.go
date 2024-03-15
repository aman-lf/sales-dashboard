package service

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type SSEController struct {
	clients map[*gin.Context]struct{}
	mutex   sync.Mutex
}

func NewSSEController() *SSEController {
	return &SSEController{
		clients: make(map[*gin.Context]struct{}),
	}
}

func (c *SSEController) AddClient(ctx *gin.Context) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clients[ctx] = struct{}{}
}

func (c *SSEController) RemoveClient(ctx *gin.Context) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.clients, ctx)
}

func (c *SSEController) TriggerEvent(eventData string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for ctx := range c.clients {
		ctx.Data(http.StatusOK, "text/event-stream", []byte("data: "+eventData+"\n\n"))
	}
}
