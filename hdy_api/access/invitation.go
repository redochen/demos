package access

import (
	"errors"
	"github.com/go-xorm/xorm"
	. "github.com/redochen/demos/hdy_api/entities"
	"github.com/redochen/demos/hdy_api/utils"
	"math"
)

//添加邀请
func AddInvitation(invitation *InvitationEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == invitation {
		return 0, errors.New("invalid parameter")
	}

	if invitation.UserId <= 0 {
		return 0, errors.New("invalid user ID")
	}

	if invitation.FriendUserId <= 0 {
		return 0, errors.New("invalid friend user ID")
	}

	return engine.InsertOne(invitation)
}

//更新邀请
func UpdateInvitation(invitation *InvitationEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == invitation {
		return 0, errors.New("invalid parameter")
	}

	if invitation.Id <= 0 {
		return 0, errors.New("invalid invitation ID")
	}

	return engine.Id(invitation.Id).Update(invitation)
}

//获取邀请
func GetInvitation(userId, friendUserId int64) (*InvitationEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if userId <= 0 {
		return nil, errors.New("invalid user ID")
	}

	if friendUserId <= 0 {
		return nil, errors.New("invalid friend user ID")
	}

	var invitation InvitationEntity

	_, err := engine.Where("user_id = ?", userId).
		And("friend_user_id = ?", friendUserId).Get(&invitation)

	return &invitation, err
}

//获取邀请列表
func GetInvitations(userId int64, status, pageIndex, pageSize int) (invitations []*InvitationEntity, totalCount, pageCount int64, err error) {
	if nil == engine {
		err = errors.New("engine not initialized")
		return
	}

	if userId <= 0 {
		err = errors.New("invalid user ID")
		return
	}

	invitation := new(InvitationEntity)
	offset, limit := utils.GetOffsetAndLimit(pageIndex, pageSize)

	totalCount, err = engine.Count(invitation)
	if err != nil {
		return
	}

	if totalCount > 0 && limit > 0 {
		pageCount = int64(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var rows *xorm.Rows

	if utils.InvitationStatusAll == status {
		rows, err = engine.Where("user_id = ?", userId).Limit(limit, offset).Rows(invitation)
	} else {
		rows, err = engine.Where("user_id = ?", userId).And("status = ?", status).Limit(limit, offset).Rows(invitation)
	}

	if err != nil {
		return
	}

	defer rows.Close()

	invitations = make([]*InvitationEntity, 0)

	for rows.Next() {
		err = rows.Scan(invitation)
		if err != nil {
			return
		}

		invitations = append(invitations, invitation)

		//注意：这里应重新分配内存
		invitation = new(InvitationEntity)
	}

	return
}

//更新邀请状态
func UpdateInvitationStatus(invitationId int64, status int) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if invitationId <= 0 {
		return 0, errors.New("invalid invitation ID")
	}

	return engine.Id(invitationId).Cols("status").Update(&InvitationEntity{Status: status})
}
