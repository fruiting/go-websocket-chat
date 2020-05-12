package app

import "github.com/gorilla/websocket"

// User structure describes user
type User struct {
	ID         string `json:"user_id"`
	IsSender   bool
	connection *websocket.Conn
}
