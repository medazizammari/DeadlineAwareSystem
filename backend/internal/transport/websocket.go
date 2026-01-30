package transport

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/medazizammari/real-time-deadline-aware-golang/internal/domain"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Handler(events <-chan domain.Event) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		log.Println("WS CONNECT", r.RemoteAddr)
		defer conn.Close()

		for event := range events {
			data, _ := json.Marshal(event)
			log.Println("WS SEND", event.ID)
			conn.WriteMessage(websocket.TextMessage, data)
		}
	}
}
