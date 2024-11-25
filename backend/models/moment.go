package models

import (
	"time"
)

type Moment struct {
	Moment_id  int
	Content    string
	User_id    string
	Created_at time.Time
	Pic_url    string
}

func (Moment) TableName() string {
	return "moment_info" // 数据库中的表名
}
