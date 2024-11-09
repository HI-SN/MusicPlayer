package models

import "time"

type Playlist struct {
	Playlist_id int
	Title       string
	User_id     string
	Create_at   time.Time
	Description string
	Type        string
	Hits        int
}

func (Playlist) TableName() string {
	return "playlist_info" // 数据库中的表名
}
