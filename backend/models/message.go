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

// 定义请求结构体，对应前端发送的JSON数据格式
type SendMessageRequest struct {
	SenderID    string `json:"sender_id"`
	ReceiverID  string `json:"receiver_id"`
	Content     string `json:"content"`
	MessageType string `json:"message_type"`
	IsRead      bool   `json:"is_read"`
}

func (Message) TableName() string {
	return "message_info"
}
