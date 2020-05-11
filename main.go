package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fruiting/go-chat/app"
)

func main() {
	fmt.Println("Listening...")

	http.HandleFunc("/ws", app.Listen)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
