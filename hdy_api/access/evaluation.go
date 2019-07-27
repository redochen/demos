package access

import (
	"errors"
	"github.com/go-xorm/xorm"
	. "github.com/redochen/demos/hdy_api/entities"
	"github.com/redochen/demos/hdy_api/utils"
	"math"
)

//添加评价
func AddEvaluation(evaluation *EvaluationEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == evaluation {
		return 0, errors.New("invalid parameter")
	}

	if evaluation.InvitationId <= 0 {
		return 0, errors.New("invalid invitation ID")
	}

	return engine.InsertOne(evaluation)
}

//更新评价
func UpdateEvaluation(evaluation *EvaluationEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == evaluation {
		return 0, errors.New("invalid parameter")
	}

	if evaluation.Id <= 0 {
		return 0, errors.New("invalid evaluation ID")
	}

	return engine.Id(evaluation.Id).Update(evaluation)
}

//获取评价列表
func GetEvaluations(invitationId int64, status, pageIndex, pageSize int) (evaluations []*EvaluationEntity, totalCount, pageCount int64, err error) {
	if nil == engine {
		err = errors.New("engine not initialized")
		return
	}

	if invitationId <= 0 {
		err = errors.New("invalid invitation ID")
		return
	}

	evaluation := new(EvaluationEntity)
	offset, limit := utils.GetOffsetAndLimit(pageIndex, pageSize)

	totalCount, err = engine.Count(evaluation)
	if err != nil {
		return
	}

	if totalCount > 0 && limit > 0 {
		pageCount = int64(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var rows *xorm.Rows
	rows, err = engine.Where("invitation_id", invitationId).Limit(limit, offset).Rows(evaluation)
	if err != nil {
		return
	}

	defer rows.Close()

	evaluations = make([]*EvaluationEntity, 0)

	for rows.Next() {
		err = rows.Scan(evaluation)
		if err != nil {
			return
		}

		evaluations = append(evaluations, evaluation)

		//注意：这里应重新分配内存
		evaluation = new(EvaluationEntity)
	}

	return
}

//删除评价
func DeleteEvaluation(evaluationId int64) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if evaluationId <= 0 {
		return 0, errors.New("invalid evaluation ID")
	}

	var evaluation EvaluationEntity
	return engine.Id(evaluationId).Delete(&evaluation)
}
