package config

import (
	"errors"

	"github.com/redochen/tools/config"
	CcFunc "github.com/redochen/tools/function"
)

// EndpointMap 服务终端定义
type EndpointMap map[string]string

var (
	//Endpoints 服务终端集合
	Endpoints EndpointMap
)

//InitEndpoints 初始化服务终端列表
func InitEndpoints() {
	options, err := config.Conf.GetOptions(sectionEndpoints)
	if err != nil {
		panic("Failed to load Endpoints config.")
	}

	Endpoints = make(map[string]string)
	for _, option := range options {
		if _, ok := Endpoints[option]; ok {
			panic("[" + option + "] already exists.")
		}

		Endpoints[option], _ = config.Conf.String(sectionEndpoints, option)
	}

	//fmt.Printf("%v\n", Endpoints)
}

//Get 根据键名获取服务终端
func (m EndpointMap) Get(key string) (string, error) {
	defer CcFunc.CheckPanic()

	endpoint, ok := m[key]
	if !ok {
		return "", errors.New("[" + key + "] not exists.")
	}

	return endpoint, nil
}
