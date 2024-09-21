package service

import (
	"Gim/internal/server"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SendMessage godoc
//	@Summary	SendMessage
//	@Tags		Message
//	@Router		/message/sendMessage [get]
func SendMessage(c *gin.Context) {
	// Upgrade http request to websocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Upgrade failed: ", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("Close connection failed: ", err)
		}
	}()
	// Create a channel to receive messages from the subscriber
	messages := make(chan string)
	ctx := c.Request.Context()
	go server.Subscribe(ctx, "websocket", messages)
	for {
		select {
		case msg := <-messages:
			// Send messages received from the subscriber to the websocket client
			time := time.Now().Format("2006-01-02 15:04:05")
			message := fmt.Sprintf("[websocket][%s]: %s", time, msg)
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				fmt.Println("Write message failed: ", err)
				return
			}
		case <-ctx.Done():
			// The client disconnected or the request was cancelled, clean up
			fmt.Println("WebSocket connection closed or request cancelled")
			return
		}
	}
}
