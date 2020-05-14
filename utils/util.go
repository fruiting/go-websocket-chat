package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON returns response with json to front
func RespondWithJSON(w http.ResponseWriter, slice []interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slice)
}
