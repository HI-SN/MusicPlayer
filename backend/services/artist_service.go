package services

import (
	"backend/database"
	"backend/models"
)

type ArtistService struct {
}

// 创建歌手信息
func (a *ArtistService) CreateArtist(artist *models.Artist) error {
	query := `INSERT INTO artist_info (name, bio, profile_pic, type, nation) VALUES (?, ?, ?, ?, ?)`
	_, err := database.DB.Exec(query, artist.Name, artist.Bio, artist.Profile_pic, artist.Type, artist.Nation)
	return err
}

// 更新歌手信息
func (a *ArtistService) UpdateArtist(artist *models.Artist) error {
	query := `UPDATE artist_info SET name = ?, bio = ?, profile_pic = ?, 
	type = ?, nation = ? WHERE id = ?`
	_, err := database.DB.Exec(query, artist.Name, artist.Bio, artist.Profile_pic, artist.Type, artist.Nation, artist.Artist_id)
	return err
}

// 通过id获取歌手信息
func (a *ArtistService) GetArtist(artistID int) (*models.Artist, error) {
	artist := &models.Artist{}
	query := "SELECT * FROM artist_info WHERE id=?"
	err := database.DB.QueryRow(query, artistID).Scan(&artist.Artist_id, &artist.Name, &artist.Bio, &artist.Profile_pic, &artist.Type, &artist.Nation)
	if err != nil {
		return nil, err
	}
	return artist, nil
}

// 删除歌手信息
func (a *ArtistService) DeleteArtist(artistID int) error {
	// 删除歌手的关联内容（专辑、歌曲等）

	// 删除歌手
	query := `DELETE FROM artist_info WHERE id = ?`
	_, err := database.DB.Exec(query, artistID)
	return err
}
