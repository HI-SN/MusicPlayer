package models

type Artist struct {
	Artist_id   int
	Name        string
	Bio         string
	Profile_pic string
	Type        string
	Nation      string
}

func (Artist) TableName() string {
	return "artist_info" // 数据库中的表名
}
