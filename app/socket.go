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

// rooms - map of rooms
var rooms map[string]Room

// reader - reads json with new message and writes it into the channel
func reader(conn *websocket.Conn, messages chan Message, r *http.Request) {
	message := Message{}
	for {
		err := conn.ReadJSON(&message)
		if err != nil {
			fmt.Println(err)
		}

		message.Room = rooms[r.URL.Query()["room"][0]]
		messages <- message
	}
}

// writer - returns message to necessary users
func writer(messages chan Message) {
	for {
		message := <-messages
		fmt.Printf("%+v\n", message)

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
func Listen(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Connected...")

	roomNumber := r.URL.Query()["room"][0]
	connectUser(websocket, roomNumber)
	messages := make(chan Message)
	go reader(websocket, messages, r)
	writer(messages)
}

func init() {
	rooms = make(map[string]Room)
}
