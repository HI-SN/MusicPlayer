package models

// FollowUser represents the follow_user table
type FollowUser struct {
	Follower_id string `json:"follower_id"`
	Followed_id string `json:"followed_id"`
}

func (FollowUser) TableName() string {
	return "follow_user"
}
