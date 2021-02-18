package access

import (
	"errors"
	"math"

	"github.com/go-xorm/xorm"
	"github.com/redochen/demos/hdy_api/entities"
	"github.com/redochen/demos/hdy_api/utils"
)

//AddGameRole 添加游戏角色
func AddGameRole(role *entities.GameRoleEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == role {
		return 0, errors.New("invalid parameter")
	}

	return engine.InsertOne(role)
}

//GetGameRoles 获取游戏角色
func GetGameRoles(userID int64, pageIndex, pageSize int) (roles []*entities.GameRoleEntity, totalCount, pageCount int64, err error) {
	if nil == engine {
		err = errors.New("engine not initialized")
		return
	}

	if userID <= 0 {
		err = errors.New("invalid user ID")
		return
	}

	role := new(entities.GameRoleEntity)
	offset, limit := utils.GetOffsetAndLimit(pageSize, pageIndex)

	totalCount, err = engine.Count(role)
	if err != nil {
		return
	}

	if totalCount > 0 && limit > 0 {
		pageCount = int64(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var rows *xorm.Rows
	rows, err = engine.Where("user_id = ?", userID).Limit(limit, offset).Rows(role)
	if err != nil {
		return
	}

	defer rows.Close()

	roles = make([]*entities.GameRoleEntity, 0)

	for rows.Next() {
		err = rows.Scan(role)
		if err != nil {
			return
		}

		roles = append(roles, role)

		//注意：这里应重新分配内存
		role = new(entities.GameRoleEntity)
	}

	return
}
