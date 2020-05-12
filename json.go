package main

import (
	"encoding/json"
	"io/ioutil"
)

//Config 配置文件模板
type Config struct {
	Port  string
	Token string
}

//dropErr 错误处理
func dropErr(e error) {
	if e != nil {
		panic(e)
	}
}

//ReadProfilePath 读取配置文件 | FilePath为文件路径
func ReadProfilePath(filePath string) *Config {
	//读取配置文件
	fileData, err := ioutil.ReadFile(filePath)
	dropErr(err)
	//解析配置文件
	configg := &Config{}
	json.Unmarshal([]byte(fileData), configg)
	dropErr(err)
	return configg
}

//ReadProfile 读取配置文件 | 默认路径为./config.json
func ReadProfile() *Config {
	config := &Config{}
	//读取配置文件
	fileData, err := ioutil.ReadFile("./config.json")
	dropErr(err)
	//解析配置文件
	json.Unmarshal([]byte(fileData), config)
	dropErr(err)
	return config
}
