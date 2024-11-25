package models

// FollowArtist represents the follow_artist table
type FollowArtist struct {
	Follower_id string
	Followed_id int
}

func (FollowArtist) TableName() string {
	return "follow_artist"
}
