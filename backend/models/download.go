package models

import "time"

// Download represents the download_info table
type Download struct {
	ID           int
	UserID       string
	SongID       int
	DownloadTime time.Time
	FileURL      string
}

func (Download) TableName() string {
	return "download_info"
}
