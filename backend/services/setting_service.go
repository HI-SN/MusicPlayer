package services

import (
	"backend/database"
	"backend/models"
)

type SettingService struct{}

// 创建新的设置
func (s *SettingService) CreateSetting(setting *models.Setting) error {
	query := `INSERT INTO setting_info (user_id) VALUES (?)`
	_, err := database.DB.Exec(query, setting.UserID)
	return err
}

// 更新设置
func (s *SettingService) UpdateSetting(setting *models.Setting) error {
	query := `UPDATE setting_info SET msg = ?, see_rank = ?, info_comment = ?, info_like = ?, info_msg = ?, info_sys = ?, service = ? WHERE user_id = ?`
	_, err := database.DB.Exec(query, setting.Msg, setting.See_rank, setting.Info_comment, setting.Info_like, setting.Info_msg, setting.Info_sys, setting.Service, setting.UserID)
	return err
}

// 删除设置
func (s *SettingService) DeleteSetting(userID string) error {
	query := `DELETE FROM setting_info WHERE user_id = ?`
	_, err := database.DB.Exec(query, userID)
	return err
}

// 获取用户的设置
func (s *SettingService) GetSetting(userID string) (*models.Setting, error) {
	setting := &models.Setting{}
	query := `SELECT user_id, msg, see_rank, info_comment, info_like, info_msg, info_sys, service FROM setting_info WHERE user_id = ?`
	err := database.DB.QueryRow(query, userID).Scan(&setting.UserID, &setting.Msg, &setting.See_rank, &setting.Info_comment, &setting.Info_like, &setting.Info_msg, &setting.Info_sys, &setting.Service)
	if err != nil {
		return nil, err
	}
	return setting, nil
}
