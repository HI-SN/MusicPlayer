package models

// Like represents the like_info table
type Like struct {
	Moment_id int    `josn:"moment_id"`
	User_id   string `json:"user_id"`
}

func (Like) TableName() string {
	return "like_info"
}
