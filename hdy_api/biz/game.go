package biz

import (
	"github.com/redochen/demos/hdy_api/access"
	"github.com/redochen/demos/hdy_api/models"
)

//GetGames 获取游戏列表
func GetGames() ([]*models.GameModel, error) {
	entities, err := access.GetGames()
	if err != nil {
		return nil, err
	}

	games := make([]*models.GameModel, 0)

	for _, entity := range entities {
		if nil == entity {
			continue
		}

		game := &models.GameModel{
			ID:   entity.ID,
			Name: entity.Name,
		}

		if game != nil {
			games = append(games, game)
		}
	}

	return games, nil
}

//GetGameAreas 获取游戏大区列表
func GetGameAreas(gameID int) ([]*models.AreaModel, error) {
	entities, err := access.GetGameAreas(gameID)
	if err != nil {
		return nil, err
	}

	gameAreas := make([]*models.AreaModel, 0)

	for _, entity := range entities {
		if nil == entity {
			continue
		}

		gameArea := &models.AreaModel{
			ID:     entity.ID,
			GameID: entity.GameID,
			Name:   entity.Name,
		}

		if gameArea != nil {
			gameAreas = append(gameAreas, gameArea)
		}
	}

	return gameAreas, nil
}

//GetGameServers 获取游戏服务器列表
func GetGameServers(gameAreaID int) ([]*models.ServerModel, error) {
	entities, err := access.GetGameServers(gameAreaID)
	if err != nil {
		return nil, err
	}

	gameServers := make([]*models.ServerModel, 0)

	for _, entity := range entities {
		if nil == entity {
			continue
		}

		gameServer := &models.ServerModel{
			ID:     entity.ID,
			AreaID: entity.GameAreaID,
			Name:   entity.Name,
		}

		if gameServer != nil {
			gameServers = append(gameServers, gameServer)
		}
	}

	return gameServers, nil
}

//GetGameDans 获取游戏段位列表
func GetGameDans(gameID int) ([]*models.DanModel, error) {
	entities, err := access.GetGameDans(gameID)
	if err != nil {
		return nil, err
	}

	gameDans := make([]*models.DanModel, 0)

	for _, entity := range entities {
		if nil == entity {
			continue
		}

		gameDan := &models.DanModel{
			ID:     entity.ID,
			GameID: entity.GameID,
			Name:   entity.Name,
		}

		if gameDan != nil {
			gameDans = append(gameDans, gameDan)
		}
	}

	return gameDans, nil
}

//GetGameHeroes 获取游戏英雄列表
func GetGameHeroes(gameID int) ([]*models.HeroModel, error) {
	entities, err := access.GetGameHeros(gameID)
	if err != nil {
		return nil, err
	}

	gameHeroes := make([]*models.HeroModel, 0)

	for _, entity := range entities {
		if nil == entity {
			continue
		}

		gameHero := &models.HeroModel{
			ID:     entity.ID,
			GameID: entity.GameID,
			Name:   entity.Name,
		}

		if gameHero != nil {
			gameHeroes = append(gameHeroes, gameHero)
		}
	}

	return gameHeroes, nil
}
