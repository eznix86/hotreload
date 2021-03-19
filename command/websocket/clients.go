package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	conn *websocket.Conn
	handler *HotReloadHandler
}

func (c *Client) watchConnection()  {
	defer func() {
		delete(c.handler.Clients, c)
		log.Println("Client disconnected !")
		log.Printf("Remaining Clients: %d", len(c.handler.Clients))
		log.Println()
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