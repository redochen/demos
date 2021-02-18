package access

import (
	"errors"
	"math"

	"github.com/go-xorm/xorm"
	"github.com/redochen/demos/hdy_api/entities"
	"github.com/redochen/demos/hdy_api/utils"
)

//AddRelation 添加关系
func AddRelation(relation *entities.RelationEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == relation {
		return 0, errors.New("invalid parameter")
	}

	if relation.UserID <= 0 {
		return 0, errors.New("invalid user ID")
	}

	if relation.FriendUserID <= 0 {
		return 0, errors.New("invalid friend user ID")
	}

	return engine.InsertOne(relation)
}

//UpdateRelation 更新关系
func UpdateRelation(relation *entities.RelationEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == relation {
		return 0, errors.New("invalid parameter")
	}

	if relation.ID <= 0 {
		return 0, errors.New("invalid relation ID")
	}

	return engine.Id(relation.ID).Update(relation)
}

//GetRelation 获取关系
func GetRelation(userID, friendUserID int64) (*entities.RelationEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}

	if friendUserID <= 0 {
		return nil, errors.New("invalid friend user ID")
	}

	var relation entities.RelationEntity

	_, err := engine.Where("user_id = ?", userID).
		And("friend_user_id = ?", friendUserID).Get(&relation)

	return &relation, err
}

//GetRelations 获取关系列表
func GetRelations(userID int64, status, pageIndex, pageSize int) (relations []*entities.RelationEntity, totalCount, pageCount int64, err error) {
	if nil == engine {
		err = errors.New("engine not initialized")
		return
	}

	if userID <= 0 {
		err = errors.New("invalid user ID")
		return
	}

	relation := new(entities.RelationEntity)
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
		rows, err = engine.Where("user_id = ?", userID).Limit(limit, offset).Rows(relation)
	} else {
		rows, err = engine.Where("user_id = ?", userID).And("status = ?", status).Limit(limit, offset).Rows(relation)
	}

	if err != nil {
		return
	}

	defer rows.Close()

	relations = make([]*entities.RelationEntity, 0)

	for rows.Next() {
		err = rows.Scan(relation)
		if err != nil {
			return
		}

		relations = append(relations, relation)

		//注意：这里应重新分配内存
		relation = new(entities.RelationEntity)
	}

	return
}

//UpdateRelationStatus 更新关系状态
func UpdateRelationStatus(relationID int64, status int) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if relationID <= 0 {
		return 0, errors.New("invalid relation ID")
	}

	return engine.Id(relationID).Cols("status").Update(&entities.RelationEntity{Status: status})
}
