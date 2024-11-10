package models

// UserLikePlaylist represents the user_like_playlist table
type UserLikePlaylist struct {
	UserID     string
	PlaylistID int
}

func (UserLikePlaylist) TableName() string {
	return "user_like_playlist"
}
