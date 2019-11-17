package handlers

import (
	"github.com/gorilla/websocket"
	"github.com/terawork-com/message-socket-service/pkg/log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return ws, err
	}
	return ws, nil
}
