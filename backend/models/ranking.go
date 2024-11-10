package models

// Ranking represents the ranking_info table
type Ranking struct {
	SongID int
	Name   string
	Rank   int
}

func (Ranking) TableName() string {
	return "ranking_info"
}
