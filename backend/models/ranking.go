package models

// Ranking represents the ranking_info table
type Ranking struct {
	SongID int
	Name   string
	Rank   int
}

// Ranking结构体用于表示排行榜概要信息
type Ranking_home struct {
	// RankingName string `json:"ranking_name"`
	CoverUrl string           `json:"cover_url"`
	Songs    []Song_rank_home `json:"songs"`
}

// Song_rank_home结构体用于表示排行榜概要中的歌曲信息
type Song_rank_home struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Song_ranking_detail 结构体用于表示排行榜详情中的歌曲信息
type Song_ranking_detail struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Duration    int    `json:"duration"`
	AlbumID     int    `json:"album_id"`
	Genre       string `json:"genre"`
	ReleaseDate string `json:"release_date"`
	SongUrl     string `json:"song_url"`
	Lyrics      string `json:"lyrics"`
	SongHit     int    `json:"song_hit"`
	Liked       string
}

func (Ranking) TableName() string {
	return "ranking_info"
}
