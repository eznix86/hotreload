package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	conn *websocket.Conn
	handler *HotReloadHandler
}

func (c *Client) watchConnection()  {
	defer func() {
		delete(c.handler.clients, c)
		fmt.Println("Client disconnected !")
		fmt.Printf("Remaining clients: %d", len(c.handler.clients))
		fmt.Println()
		c.conn.Close()
	}()

	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}