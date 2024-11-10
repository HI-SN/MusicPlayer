package models

// ArtistAlbumRelation represents the artist_album_relation table
type ArtistAlbumRelation struct {
	ArtistID int
	AlbumID  int
}

func (ArtistAlbumRelation) TableName() string {
	return "artist_album_relation"
}
