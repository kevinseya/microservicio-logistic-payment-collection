package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

var wsConnections = make(map[string]*websocket.Conn)

func NotifyMicroservices(orderID, status string) {
	message := fmt.Sprintf("Payment for order %s has changed to %s", orderID, status)
	for _, conn := range wsConnections {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("WebSocket error:", err)
		}
	}
}
