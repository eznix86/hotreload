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
	Clients    map[*Client] bool
	ServerPort int
}

func (h *HotReloadHandler) Reload()  {
	for client, _ := range h.Clients {
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
		h.Clients[&c] = true
		go c.watchConnection()
		log.Println()
		log.Println("New Client !")
		log.Printf("Total Clients: %d", len(h.Clients))
		log.Println()
	}
}

func (h *HotReloadHandler) Serve() {
	h.Clients = make(map[*Client]bool)
	port := fmt.Sprintf(":%d", h.ServerPort)
	log.Println(fmt.Sprintf("Listening to port 0.0.0.0%s", port))
	if err := http.ListenAndServe(port, h); err != nil {
		log.Fatal(err)
	}
}
