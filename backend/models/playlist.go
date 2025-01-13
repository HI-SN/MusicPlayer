package models

import "time"

type Playlist struct {
	Playlist_id int       `json:"id"`
	Title       string    `json:"title"`
	User_id     string    `json:"user_id"`
	Create_at   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Hits        int       `json:"hits"`
	Cover_url   string    `json:"cover_url"`
}

// PlaylistResponse 用于首页展示歌单时返回给客户端的结构体
type PlaylistResponse struct {
	Playlist_id int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Hits        int    `json:"hits"`
	CoverUrl    string `json:"cover_url"`
}

func (Playlist) TableName() string {
	return "playlist_info" // 数据库中的表名
}
