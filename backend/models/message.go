package models

import "time"

// Message represents the message_info table
type Message struct {
	ID         int
	CreatedAt  time.Time
	SenderID   string
	ReceiverID string
	Content    string
	IsRead     bool
}

func (Message) TableName() string {
	return "message_info"
}
