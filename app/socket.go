package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrader - websocket settings
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// users - map of [user_id => websocket.conn]
var users map[string]*websocket.Conn

// reader - reads json with new message and writes it into the channel
func reader(conn *websocket.Conn, messages chan Message) {
	message := Message{}
	for {
		err := conn.ReadJSON(&message)
		if err != nil {
			fmt.Println(err)
		}

		message.User.connection = users[message.User.ID]
		messages <- message
	}
}

func writer(messages chan Message) {
	for {
		message := <-messages
		for _, user := range users {
			user.WriteJSON(message)
		}
	}
}

// connectUser creates User object with user id and connection and appends it into all users slice
func connectUser(connection *websocket.Conn) {
	if users == nil {
		users = make(map[string]*websocket.Conn)
	}

	user := User{}
	connection.ReadJSON(&user)
	users[user.ID] = connection
}

// Listen function opens connection for users and listens any new messages from them
func Listen(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	connectUser(websocket)
	messages := make(chan Message)
	go reader(websocket, messages)
	writer(messages)
}
