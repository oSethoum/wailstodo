package handlers

import (
	"github.com/gofiber/websocket/v2"
)

func Subscribe(c *websocket.Conn) {
	channel := make(chan string, 1)
	Subs[&channel] = channel

	for {
		c.WriteMessage(websocket.TextMessage, []byte(<-Subs[&channel]))
	}
}
