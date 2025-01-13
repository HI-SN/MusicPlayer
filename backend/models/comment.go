package models

import (
	"time"
)

type Comment struct {
	Comment_id int       `json:"comment_id"`
	Content    string    `json:"content"`
	User_id    string    `json:"user_id"`
	User_name  string    `json:"user_name"`
	Created_at time.Time `json:"created_at"`
	Type       string    `json:"type"`
	Target_id  int       `json:"target_id"`
}

// 用于返回给前端的结构体
type MomentComment struct {
	Comment_id int       `json:"comment_id"`
	Content    string    `json:"content"`
	User_id    string    `json:"user_id"`
	User_name  string    `json:"user_name"`
	Created_at time.Time `json:"created_at"`
}

func (Comment) TableName() string {
	return "comment_info" // 数据库中的表名
}
