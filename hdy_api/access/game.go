package access

import (
	"errors"

	"github.com/redochen/demos/hdy_api/entities"
)

//GetGames 获取游戏列表
func GetGames() ([]*entities.GameEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	game := new(entities.GameEntity)
	games := make([]*entities.GameEntity, 0)

	rows, err := engine.Rows(game)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(game)
		if err != nil {
			return nil, err
		}

		games = append(games, game)

		//注意：这里应重新分配内存
		game = new(entities.GameEntity)
	}

	return games, nil
}

//GetGameAreas 获取游戏大区列表
func GetGameAreas(gameID int) ([]*entities.GameAreaEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if gameID <= 0 {
		return nil, errors.New("invalid game ID")
	}

	area := new(entities.GameAreaEntity)
	areas := make([]*entities.GameAreaEntity, 0)

	rows, err := engine.Where("game_id = ?", gameID).Rows(area)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(area)
		if err != nil {
			return nil, err
		}

		areas = append(areas, area)

		//注意：这里应重新分配内存
		area = new(entities.GameAreaEntity)
	}

	return areas, nil
}

//GetGameServers 获取游戏服务器列表
func GetGameServers(gameAreaID int) ([]*entities.GameServerEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if gameAreaID <= 0 {
		return nil, errors.New("invalid game area ID")
	}

	server := new(entities.GameServerEntity)
	servers := make([]*entities.GameServerEntity, 0)

	rows, err := engine.Where("game_area_id = ?", gameAreaID).Rows(server)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(server)
		if err != nil {
			return nil, err
		}

		servers = append(servers, server)

		//注意：这里应重新分配内存
		server = new(entities.GameServerEntity)
	}

	return servers, nil
}

//GetGameDans 获取游戏段位列表
func GetGameDans(gameID int) ([]*entities.GameDanEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if gameID <= 0 {
		return nil, errors.New("invalid game ID")
	}

	dan := new(entities.GameDanEntity)
	dans := make([]*entities.GameDanEntity, 0)

	rows, err := engine.Where("game_id = ?", gameID).Rows(dan)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(dan)
		if err != nil {
			return nil, err
		}

		dans = append(dans, dan)

		//注意：这里应重新分配内存
		dan = new(entities.GameDanEntity)
	}

	return dans, nil
}

//GetGameHeros 获取游戏英雄列表
func GetGameHeros(gameID int) ([]*entities.GameHeroEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if gameID <= 0 {
		return nil, errors.New("invalid game ID")
	}

	hero := new(entities.GameHeroEntity)
	heros := make([]*entities.GameHeroEntity, 0)

	rows, err := engine.Where("game_id = ?", gameID).Rows(hero)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(hero)
		if err != nil {
			return nil, err
		}

		heros = append(heros, hero)

		//注意：这里应重新分配内存
		hero = new(entities.GameHeroEntity)
	}

	return heros, nil
}
