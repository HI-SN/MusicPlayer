package models

// ArtistSongRelation represents the artist_song_relation table
type ArtistSongRelation struct {
	ArtistID int
	SongID   int
}

func (ArtistSongRelation) TableName() string {
	return "artist_song_relation"
}
