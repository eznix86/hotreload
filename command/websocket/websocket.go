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
	}
)

type HotReloadHandler struct {
	reload chan bool
}

func (h *HotReloadHandler) Reload()  {
	h.reload <- true
}

func (h *HotReloadHandler) reader(ws *websocket.Conn) {
	defer ws.Close()
	for {
		messageType, p, err := ws.ReadMessage()
		fmt.Println(err)
		fmt.Println(messageType)
		fmt.Println(string(p))
		fmt.Println()
		if err != nil {
			break
		}
	}
}

func (h *HotReloadHandler) writer(ws *websocket.Conn) {
	for {
		select {
		case <- h.reload:
			var p []byte
			if err := ws.WriteMessage(websocket.TextMessage, p); err != nil {
				return
			}
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

		go h.writer(ws)
		h.reader(ws)
	}
}

func (h *HotReloadHandler) Serve() {
	h.reload = make(chan bool)
	if err := http.ListenAndServe(":9023", h); err != nil {
		log.Fatal(err)
	}
}
