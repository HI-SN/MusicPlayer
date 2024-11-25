package services

import (
	"backend/database"
	"backend/models"
	"io"
	"os"
	"path/filepath"
	"time"
)

// UploadService 定义上传相关的服务函数
type UploadService struct{}

// UploadAudio 上传音频文件
func (u *UploadService) UploadAudio(file *os.File, filename string) (int, error) {
	// 保存文件到指定目录
	uploadPath := "./uploads/" + filename
	err := os.MkdirAll(filepath.Dir(uploadPath), os.ModePerm)
	if err != nil {
		return 0, err
	}

	// 将文件保存到指定目录
	out, err := os.Create(uploadPath)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	_, err = file.Seek(0, 0)
	if err != nil {
		return 0, err
	}

	_, err = io.Copy(out, file)
	if err != nil {
		return 0, err
	}

	// 将文件名存储到数据库中，并返回生成的歌曲ID
	query := "INSERT INTO song_info (song_url, created_at, updated_at) VALUES ($1, $2, $3) RETURNING song_id"
	var songID int
	err = database.DB.QueryRow(query, filename, time.Now(), time.Now()).Scan(&songID)
	if err != nil {
		return 0, err
	}
	return songID, nil
}

// UploadSongInfo 上传歌曲信息
func (u *UploadService) UploadSongInfo(title string, duration int, albumID int, genre string, releaseDate time.Time, songURL string, lyrics string, songHit int) (int, error) {
	song := &models.Song{
		Title:        title,
		Duration:     duration,
		Album_id:     albumID,
		Genre:        genre,
		Release_date: releaseDate,
		Song_url:     songURL,
		Lyrics:       lyrics,
		Song_hit:     songHit,
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
	}
	return song.Song_id, (&SongService{}).CreateSong(song)
}

// UploadLyrics 上传歌词文件
func (u *UploadService) UploadLyrics(songID int, file *os.File, filename string) error {
	// 保存文件到指定目录
	uploadPath := "./uploads/" + filename
	err := os.MkdirAll(filepath.Dir(uploadPath), os.ModePerm)
	if err != nil {
		return err
	}

	// 将文件保存到指定目录
	out, err := os.Create(uploadPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	// 将文件名存储到数据库中，并更新歌曲的歌词字段
	query := "UPDATE song_info SET lyrics = $1 WHERE song_id = $2"
	_, err = database.DB.Exec(query, filename, songID)
	return err
}
