package models

// Setting represents the setting_info table
type Setting struct {
	UserID       string `json:"user_id"`
	Msg          int    `json:"msg"`
	See_rank     int    `json:"rank"`
	Info_comment int    `json:"info_comment"`
	Info_like    int    `json:"info_like"`
	Info_msg     int    `json:"info_msg"`
	Info_sys     int    `json:"info_sys"`
	Service      int    `json:"service"`
}

func (Setting) TableName() string {
	return "setting_info"
}
