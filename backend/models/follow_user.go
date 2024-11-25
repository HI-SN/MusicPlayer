package models

// FollowUser represents the follow_user table
type FollowUser struct {
	Follower_id string
	Followed_id string
}

func (FollowUser) TableName() string {
	return "follow_user"
}
