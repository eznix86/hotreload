package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
var (
	upgrader  = websocket.Upgrader {
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

type HotReloadHandler struct {
	clients map[*Client] bool
}

func (h *HotReloadHandler) Reload()  {
	for client, _ := range h.clients {
		var p []byte
		if err := client.conn.WriteMessage(websocket.TextMessage, p); err != nil {
		   return
		}
	}
}

func (h *HotReloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/ws" {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			if _, ok := err.(websocket.HandshakeError); !ok {
				log.Println(err)
			}
			return
		}

		c := Client{conn: ws, handler: h}
		h.clients[&c] = true
		go c.watchConnection()
		fmt.Println()
		fmt.Println("New Client !")
		fmt.Printf("Total clients: %d", len(h.clients))
		fmt.Println()
	}
}

func (h *HotReloadHandler) Serve() {
	h.clients = make(map[*Client]bool)
	if err := http.ListenAndServe(":9023", h); err != nil {
		log.Fatal(err)
	}
}
