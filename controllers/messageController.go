package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fruiting/go-chat/app"
)

// GetMessages function returns slice of
func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(app.GetMessagesSlice("1", 10))
}
