package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	WatchTV struct {
		HTTP string `yaml:"http"`
	} `yaml:"watchtv"`
	Databases struct {
		MariaDB struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
			Charset  string `yaml:"charset"`
		} `yaml:"mariadb"`
	} `yaml:"databases"`
}

var Config ConfigStruct

func init() {
	loadConfig()
	mariaDBConnect()
}

func loadConfig() {
	// 读取 Yaml 文件
	data, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	// 解析 Yaml 数据
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}
}
