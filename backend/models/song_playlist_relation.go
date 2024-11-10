package models

// SongPlaylistRelation represents the song_playlist_relation table
type SongPlaylistRelation struct {
	PlaylistID int
	SongID     int
}

func (SongPlaylistRelation) TableName() string {
	return "song_playlist_relation"
}
