package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	WatchTV struct {
		HTTP string `yaml:"http"`
	} `yaml:"watchtv"`
}

var Config ConfigStruct

func init() {
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
