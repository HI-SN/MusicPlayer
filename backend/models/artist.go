package models

import "database/sql"

type Artist struct {
	Artist_id   int
	Name        string
	Bio         string
	Profile_pic string
	Type        string
	Nation      string
}

func (Artist) TableName() string {
	return "artist_info" // 数据库中的表名
}

// 创建歌手信息
func CreateArtist(db *sql.DB, a *Artist) error {
	query := `INSERT INTO artist_info (name, bio, profile_pic, type, nation) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, a.Name, a.Bio, a.Profile_pic, a.Type, a.Nation)
	return err
}

// 更新歌手信息
func UpdateArtist(db *sql.DB, a *Artist) error {
	query := `UPDATE artist_info SET name = ?, bio = ?, profile_pic = ?, 
	type = ?, nation = ? WHERE id = ?`
	_, err := db.Exec(query, a.Name, a.Bio, a.Profile_pic, a.Type, a.Nation, a.Artist_id)
	return err
}

// 通过id获取歌手信息
func GetArtist(db *sql.DB, artistID int) (*Artist, error) {
	a := &Artist{}
	query := "SELECT * FROM artist_info WHERE id=?"
	err := db.QueryRow(query, artistID).Scan(&a.Artist_id, &a.Name, &a.Bio, &a.Profile_pic, &a.Type, &a.Nation)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// 删除歌手信息
func DeleteArtist(db *sql.DB, artistID int) error {
	// 删除歌手的关联内容（专辑、歌曲等）

	// 删除歌手
	query := `DELETE FROM artist_info WHERE id = ?`
	_, err := db.Exec(query, artistID)
	return err
}
