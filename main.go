package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fruiting/go-chat/app"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	fmt.Println("Listening...")

	http.HandleFunc("/chat", app.Listen)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("LISTENER_PORT"), nil))
}
