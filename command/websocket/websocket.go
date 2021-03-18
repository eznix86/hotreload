package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
var (
	upgrader  = websocket.Upgrader {
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type HotReloadHandler struct {
	clients []*websocket.Conn
}

func (h *HotReloadHandler) Reload()  {
	for _, ws := range h.clients {
		var p []byte
		if err := ws.WriteMessage(websocket.TextMessage, p); err != nil {
		   return
		}
	}
}

func (h *HotReloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/ws" {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			if _, ok := err.(websocket.HandshakeError); !ok {
				log.Println(err)
			}
			return
		}

		h.clients = append(h.clients, ws)
	}
}

func (h *HotReloadHandler) Serve() {
	if err := http.ListenAndServe(":9023", h); err != nil {
		log.Fatal(err)
	}
}
