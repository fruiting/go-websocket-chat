package app

import (
	"net/http"

	"github.com/fruiting/go-chat/logger"
	"github.com/gorilla/websocket"
)

// rooms - map of rooms
var rooms map[string]Room

// reader - reads json with new message and writes it into the channel
func reader(conn *websocket.Conn, messages chan Message, roomNumber string) {
	message := Message{}
	for {
		err := conn.ReadJSON(&message)
		if err != nil {
			logger.Error("Unable to read message in room " + roomNumber + ". Reaseon: " + err.Error())
		}

		message.Room = rooms[roomNumber]
		message.Room.ID = roomNumber
		messages <- message
	}
}

// writer - returns message to necessary users
func writer(messages chan Message) {
	for {
		message := <-messages

		SaveMessage(message)
		for _, user := range message.Room.Users {
			user.connection.WriteJSON(message)
		}
	}
}

// connectUser creates User object with user id and connection and appends it into all users slice
func connectUser(connection *websocket.Conn, roomNumber string) {
	user := User{}
	connection.ReadJSON(&user)
	user.connection = connection

	room := Room{roomNumber, []User{}}
	if rooms[roomNumber].ID == "" {
		room = rooms[roomNumber]
		room.Users = append(room.Users, user)
	} else {
		room = Room{roomNumber, []User{user}}
	}
	rooms[roomNumber] = room
}

// Listen function opens connection for users and listens any new messages from them
func Listen(upgrader websocket.Upgrader, roomNumber string, w http.ResponseWriter, r *http.Request) {
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("Unable to connect. Error: " + err.Error())
	}

	connectUser(websocket, roomNumber)
	messages := make(chan Message)
	go reader(websocket, messages, roomNumber)
	writer(messages)
}

func init() {
	rooms = make(map[string]Room)
}
