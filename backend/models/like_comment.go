package models

// LikeComment represents the like_comment table
type LikeComment struct {
	Comment_id int    `josn:"comment_id"`
	User_id    string `json:"user_id"`
}

func (LikeComment) TableName() string {
	return "like_comment"
}
