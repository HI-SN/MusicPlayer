package models

// Like represents the like_info table
type Like struct {
	MomentID int
	UserID   string
}

func (Like) TableName() string {
	return "like_info"
}
