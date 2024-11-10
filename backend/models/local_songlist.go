package models

import "time"

// LocalSonglist represents the local_songlist table
type LocalSonglist struct {
	SongID    int
	UserID    string
	FileURL   string
	AddedTime time.Time
}

func (LocalSonglist) TableName() string {
	return "local_songlist"
}
