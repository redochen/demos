package access

import (
	"errors"
	"github.com/go-xorm/xorm"
	. "github.com/redochen/demos/hdy_api/entities"
	"github.com/redochen/demos/hdy_api/utils"
	"math"
)

//添加关系
func AddRelation(relation *RelationEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == relation {
		return 0, errors.New("invalid parameter")
	}

	if relation.UserId <= 0 {
		return 0, errors.New("invalid user ID")
	}

	if relation.FriendUserId <= 0 {
		return 0, errors.New("invalid friend user ID")
	}

	return engine.InsertOne(relation)
}

//更新关系
func UpdateRelation(relation *RelationEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == relation {
		return 0, errors.New("invalid parameter")
	}

	if relation.Id <= 0 {
		return 0, errors.New("invalid relation ID")
	}

	return engine.Id(relation.Id).Update(relation)
}

//获取关系
func GetRelation(userId, friendUserId int64) (*RelationEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if userId <= 0 {
		return nil, errors.New("invalid user ID")
	}

	if friendUserId <= 0 {
		return nil, errors.New("invalid friend user ID")
	}

	var relation RelationEntity

	_, err := engine.Where("user_id = ?", userId).
		And("friend_user_id = ?", friendUserId).Get(&relation)

	return &relation, err
}

//获取关系列表
func GetRelations(userId int64, status, pageIndex, pageSize int) (relations []*RelationEntity, totalCount, pageCount int64, err error) {
	if nil == engine {
		err = errors.New("engine not initialized")
		return
	}

	if userId <= 0 {
		err = errors.New("invalid user ID")
		return
	}

	relation := new(RelationEntity)
	offset, limit := utils.GetOffsetAndLimit(pageIndex, pageSize)

	totalCount, err = engine.Count(relation)
	if err != nil {
		return
	}

	if totalCount > 0 && limit > 0 {
		pageCount = int64(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var rows *xorm.Rows

	if utils.RelationStatusAll == status {
		rows, err = engine.Where("user_id = ?", userId).Limit(limit, offset).Rows(relation)
	} else {
		rows, err = engine.Where("user_id = ?", userId).And("status = ?", status).Limit(limit, offset).Rows(relation)
	}

	if err != nil {
		return
	}

	defer rows.Close()

	relations = make([]*RelationEntity, 0)

	for rows.Next() {
		err = rows.Scan(relation)
		if err != nil {
			return
		}

		relations = append(relations, relation)

		//注意：这里应重新分配内存
		relation = new(RelationEntity)
	}

	return
}

//更新关系状态
func UpdateRelationStatus(relationId int64, status int) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if relationId <= 0 {
		return 0, errors.New("invalid relation ID")
	}

	return engine.Id(relationId).Cols("status").Update(&RelationEntity{Status: status})
}
