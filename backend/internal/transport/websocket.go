package transport

import (
	"encoding/json"
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
		defer conn.Close()

		for event := range events {
			data, _ := json.Marshal(event)
			conn.WriteMessage(websocket.TextMessage, data)
		}
	}
}
