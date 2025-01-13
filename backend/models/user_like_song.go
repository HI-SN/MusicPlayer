package models

// UserLikeSong represents the user_like_song table
type UserLikeSong struct {
	UserID string `json:"user_id"`
	SongID int    `json:"song_id"`
}

type NewUserLikeSong struct {
	Song
	Artist_ids   []int    `json:"artist_id"`
	Artist_names []string `json:"artist_name"`
	Album_name   string   `json:"album_name"`
}

func (UserLikeSong) TableName() string {
	return "user_like_song"
}
