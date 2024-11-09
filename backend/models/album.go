package models

import "time"

type Album struct {
	Album_id     int
	Name         string
	Description  string
	Release_date time.Time
	Cover_url    string
}

func (Album) TableName() string {
	return "album_info" // 数据库中的表名
}
