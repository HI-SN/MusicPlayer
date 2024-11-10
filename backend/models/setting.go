package models

// Setting represents the setting_info table
type Setting struct {
	UserID string
	Type   string
	Value  string
}

func (Setting) TableName() string {
	return "setting_info"
}
