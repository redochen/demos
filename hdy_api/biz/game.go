package biz

import (
	"github.com/redochen/demos/hdy_api/access"
	. "github.com/redochen/demos/hdy_api/models"
)

//获取游戏列表
func GetGames() ([]*GameModel, error) {
	entities, err := access.GetGames()
	if err != nil {
		return nil, err
	}

	models := make([]*GameModel, 0)

	for _, entity := range entities {
		if nil == entity {
			continue
		}

		model := &GameModel{
			Id:   entity.Id,
			Name: entity.Name,
		}

		if model != nil {
			models = append(models, model)
		}
	}

	return models, nil
}

//获取游戏大区列表
func GetGameAreas(gameId int) ([]*AreaModel, error) {
	entities, err := access.GetGameAreas(gameId)
	if err != nil {
		return nil, err
	}

	models := make([]*AreaModel, 0)

	for _, entity := range entities {
		if nil == entity {
			continue
		}

		model := &AreaModel{
			Id:     entity.Id,
			GameId: entity.GameId,
			Name:   entity.Name,
		}

		if model != nil {
			models = append(models, model)
		}
	}

	return models, nil
}

//获取游戏服务器列表
func GetGameServers(gameAreaId int) ([]*ServerModel, error) {
	entities, err := access.GetGameServers(gameAreaId)
	if err != nil {
		return nil, err
	}

	models := make([]*ServerModel, 0)

	for _, entity := range entities {
		if nil == entity {
			continue
		}

		model := &ServerModel{
			Id:     entity.Id,
			AreaId: entity.GameAreaId,
			Name:   entity.Name,
		}

		if model != nil {
			models = append(models, model)
		}
	}

	return models, nil
}

//获取游戏段位列表
func GetGameDans(gameId int) ([]*DanModel, error) {
	entities, err := access.GetGameDans(gameId)
	if err != nil {
		return nil, err
	}

	models := make([]*DanModel, 0)

	for _, entity := range entities {
		if nil == entity {
			continue
		}

		model := &DanModel{
			Id:     entity.Id,
			GameId: entity.GameId,
			Name:   entity.Name,
		}

		if model != nil {
			models = append(models, model)
		}
	}

	return models, nil
}

//获取游戏英雄列表
func GetGameHeroes(gameId int) ([]*HeroModel, error) {
	entities, err := access.GetGameHeros(gameId)
	if err != nil {
		return nil, err
	}

	models := make([]*HeroModel, 0)

	for _, entity := range entities {
		if nil == entity {
			continue
		}

		model := &HeroModel{
			Id:     entity.Id,
			GameId: entity.GameId,
			Name:   entity.Name,
		}

		if model != nil {
			models = append(models, model)
		}
	}

	return models, nil
}
