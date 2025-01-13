package models

// FollowArtist represents the follow_artist table
type FollowArtist struct {
	Follower_id string `json:"follower_id"`
	Followed_id int    `json:"followed_id"`
}

func (FollowArtist) TableName() string {
	return "follow_artist"
}
