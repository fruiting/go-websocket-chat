package app

// Message structure
type Message struct {
	User     User
	Body     string `json:"body"`
	DateTime string `json:"date_time"`
}
