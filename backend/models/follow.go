package models

// Follow represents the follow_info table
type Follow struct {
	FollowerID string
	Type       string
	FollowedID string
}

func (Follow) TableName() string {
	return "follow_info"
}
