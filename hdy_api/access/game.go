package access

import (
	"errors"
	. "github.com/redochen/demos/hdy_api/entities"
)

//获取游戏列表
func GetGames() ([]*GameEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	game := new(GameEntity)
	games := make([]*GameEntity, 0)

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
		game = new(GameEntity)
	}

	return games, nil
}

//获取游戏大区列表
func GetGameAreas(gameId int) ([]*GameAreaEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if gameId <= 0 {
		return nil, errors.New("invalid game ID")
	}

	area := new(GameAreaEntity)
	areas := make([]*GameAreaEntity, 0)

	rows, err := engine.Where("game_id = ?", gameId).Rows(area)
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
		area = new(GameAreaEntity)
	}

	return areas, nil
}

//获取游戏服务器列表
func GetGameServers(gameAreaId int) ([]*GameServerEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if gameAreaId <= 0 {
		return nil, errors.New("invalid game area ID")
	}

	server := new(GameServerEntity)
	servers := make([]*GameServerEntity, 0)

	rows, err := engine.Where("game_area_id = ?", gameAreaId).Rows(server)
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
		server = new(GameServerEntity)
	}

	return servers, nil
}

//获取游戏段位列表
func GetGameDans(gameId int) ([]*GameDanEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if gameId <= 0 {
		return nil, errors.New("invalid game ID")
	}

	dan := new(GameDanEntity)
	dans := make([]*GameDanEntity, 0)

	rows, err := engine.Where("game_id = ?", gameId).Rows(dan)
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
		dan = new(GameDanEntity)
	}

	return dans, nil
}

//获取游戏英雄列表
func GetGameHeros(gameId int) ([]*GameHeroEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if gameId <= 0 {
		return nil, errors.New("invalid game ID")
	}

	hero := new(GameHeroEntity)
	heros := make([]*GameHeroEntity, 0)

	rows, err := engine.Where("game_id = ?", gameId).Rows(hero)
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
		hero = new(GameHeroEntity)
	}

	return heros, nil
}
