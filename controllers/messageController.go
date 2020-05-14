package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fruiting/go-chat/app"
)

// GetMessages function returns slice of
func GetMessages(w http.ResponseWriter, r *http.Request) {
	roomNumber := r.URL.Query()["room"][0]
	messages := app.GetMessagesSlice(roomNumber, 10)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
