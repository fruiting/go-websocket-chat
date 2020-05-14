package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fruiting/go-chat/controllers"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	fmt.Println("Listening...")

	http.HandleFunc("/api/chat", controllers.Chat)
	http.HandleFunc("/api/messages/get", controllers.GetMessages)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("LISTENER_PORT"), nil))
}
