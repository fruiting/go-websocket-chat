package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fruiting/go-chat/app"
)

func main() {
	fmt.Println("Listening...")

	http.HandleFunc("/chat", app.Listen)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
