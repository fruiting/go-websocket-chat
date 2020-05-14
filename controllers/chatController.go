package controllers

import (
	"net/http"

	"github.com/fruiting/go-chat/app"
	"github.com/gorilla/websocket"
)

// upgrader - websocket settings
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Chat function opens connection for users and listens any new messages from them
func Chat(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	roomNumber := r.URL.Query()["room"][0]
	app.Listen(upgrader, roomNumber, w, r)
}
