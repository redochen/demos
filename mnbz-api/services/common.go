package services

import (
	"github.com/redochen/demos/mnbz-api/models"
	"github.com/redochen/demos/mnbz-api/status"
)

// handleResult处理结果
func handleResult(err error, c chan<- interface{}) {
	if err != nil {
		c <- models.NewResult(status.CustomError, err.Error())
	} else {
		c <- models.NewResult(status.Success)
	}
}

// handleDataResult 处理数据结果
func handleDataResult(data interface{}, err error, c chan<- interface{}) {
	if err != nil {
		c <- models.NewResult(status.CustomError, err.Error())
	} else {
		c <- models.NewDataResult(data)
	}
}

// handleListResult 处理列表结果
func handleListResult(items interface{}, totalCount, pageCount int64, err error, c chan<- interface{}) {
	if err != nil {
		c <- models.NewResult(status.CustomError, err.Error())
	} else {
		c <- models.NewListResultEx(items, totalCount, pageCount)
	}
}
