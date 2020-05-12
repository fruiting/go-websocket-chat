package app

// Message structure
type Message struct {
	Room     Room
	Body     string `json:"body"`
	DateTime string `json:"date_time"`
}
