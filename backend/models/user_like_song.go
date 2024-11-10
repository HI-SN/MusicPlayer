package models

// UserLikeSong represents the user_like_song table
type UserLikeSong struct {
	UserID string
	SongID int
}

func (UserLikeSong) TableName() string {
	return "user_like_song"
}
