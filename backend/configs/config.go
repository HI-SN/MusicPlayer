package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Database struct {
		Type     string `json:"type"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
	Server struct {
		Port     string `json:"port"`
		LogLevel string `json:"log_level"`
	} `json:"server"`
}

var AppConfig Config

// LoadConfig 从配置文件加载配置
func LoadConfig(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	err = json.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}
}
