package models

import (
	"time"
)

type Moment struct {
	Moment_id  int       `json:"moment_id"`
	Content    string    `json:"content"`
	User_id    string    `json:"user_id"`
	Created_at time.Time `json:"created_at"`
	Pic_url    string    `json:"pic_url"`
}

type MomentAndLike struct {
	Moment
	IsLiked   bool `json:"isLiked"`
	LikeCount int  `json:"likeCount"`
}

func (Moment) TableName() string {
	return "moment_info" // 数据库中的表名
}
