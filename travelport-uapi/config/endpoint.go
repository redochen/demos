package config

import (
	"errors"
	. "github.com/redochen/tools/config"
	. "github.com/redochen/tools/function"
)

type EndpointMap map[string]string

var (
	Endpoints EndpointMap
)

func InitEndpoints() {
	options, err := Conf.GetOptions(sectionEndpoints)
	if err != nil {
		panic("Failed to load Endpoints config.")
	}

	Endpoints = make(map[string]string)
	for _, option := range options {
		if _, ok := Endpoints[option]; ok {
			panic("[" + option + "] already exists.")
		}

		Endpoints[option], _ = Conf.String(sectionEndpoints, option)
	}

	//fmt.Printf("%v\n", Endpoints)
}

func (m EndpointMap) Get(key string) (string, error) {
	defer CheckPanic()

	endpoint, ok := m[key]
	if !ok {
		return "", errors.New("[" + key + "] not exists.")
	}

	return endpoint, nil
}
