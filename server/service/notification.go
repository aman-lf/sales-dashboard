package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	messageChannel = make(chan string)
)

func SetupSSEMessage(c *gin.Context) {
	flusher, _ := c.Writer.(http.Flusher)

	// Send an initial empty message to connect client
	fmt.Fprintf(c.Writer, "data: %s\n\n", `{"message": ""}`)
	flusher.Flush()

	for {
		select {
		case message := <-messageChannel:
			fmt.Fprintf(c.Writer, "data: %s\n\n", `{"message": "`+message+`"}`)
			flusher.Flush()
		case <-c.Writer.CloseNotify():
			// Exit loop if client disconnected
			return
		}
	}
}

func TriggerSSEMessages(message string) {
	select {
	case messageChannel <- message:
	default:
	}
}
