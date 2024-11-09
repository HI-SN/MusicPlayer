package models

import "time"

type Comment struct {
	Comment_id int
	Content    string
	User_id    string
	Create_at  time.Time
	Type       string
	Target_id  int
}

func (Comment) TableName() string {
	return "comment_info" // 数据库中的表名
}
