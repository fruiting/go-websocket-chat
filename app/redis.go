package app

import (
	"encoding/json"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
)

// client *redis.Client Redis connection object
var client *redis.Client

// SaveMessage function writes Message structure to redis as json
func SaveMessage(message Message) {
	messageJSON, _ := json.Marshal(message)
	client.LPush("messages:"+message.Room.ID, string(messageJSON))
}

// GetMessages returns messages for room
// @param roomID string Number of private chat room to get access to redis
// @param count int64 Count of messages to return
// @return []*Message Slice of messages
func GetMessages(roomID string, count int64) []Message {
	messagesSlice := make([]Message, count)

	messages, _ := client.LRange("messages:"+roomID, 0, count-1).Result()
	for _, messageJSON := range messages {
		b := []byte(messageJSON)
		message := &Message{}
		json.Unmarshal(b, message)

		messagesSlice = append(messagesSlice, *message)
	}

	return messagesSlice
}

func init() {
	godotenv.Load()
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}
