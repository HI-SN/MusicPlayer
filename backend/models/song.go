package models

import (
	"time"
)

type Song struct {
	Song_id      int       `json:"id"`
	Title        string    `json:"title"`
	ArtistID     int       `json:"artist_id"`
	Duration     int       `json:"duration"`
	Album_id     int       `json:"album_id"`
	Genre        string    `json:"genre"`
	Release_date time.Time `json:"release_date"`
	Song_url     string    `json:"song_url"`
	Lyrics       string    `json:"lyrics"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	Song_hit     int       `json:"song_hit"`
}

func (Song) TableName() string {
	return "song_info"
}
