package main

import (
	"encoding/json"
	"io/ioutil"
)

//config 配置文件模板
type config struct {
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
func ReadProfilePath(filePath string) *config {
	//读取配置文件
	fileData, err := ioutil.ReadFile(filePath)
	dropErr(err)
	//解析配置文件
	config := &config{}
	json.Unmarshal([]byte(fileData), config)
	dropErr(err)
	return config
}

//ReadProfile 读取配置文件 | 默认路径为./config.json
func ReadProfile() *config {
	config := &config{}
	//读取配置文件
	fileData, err := ioutil.ReadFile("./config.json")
	dropErr(err)
	//解析配置文件
	json.Unmarshal([]byte(fileData), config)
	dropErr(err)
	return config
}
