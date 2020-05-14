package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fruiting/go-chat/controllers"
	"github.com/fruiting/go-chat/middlewares"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/chat", controllers.Chat)
	router.HandleFunc("/api/messages/get", controllers.GetMessages).Methods("GET")

	router.Use(middlewares.HeaderMiddleware)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("LISTENER_PORT"), router))
}
