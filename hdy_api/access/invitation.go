package access

import (
	"errors"
	"math"

	"github.com/go-xorm/xorm"
	"github.com/redochen/demos/hdy_api/entities"
	"github.com/redochen/demos/hdy_api/utils"
)

//AddInvitation 添加邀请
func AddInvitation(invitation *entities.InvitationEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == invitation {
		return 0, errors.New("invalid parameter")
	}

	if invitation.UserID <= 0 {
		return 0, errors.New("invalid user ID")
	}

	if invitation.FriendUserID <= 0 {
		return 0, errors.New("invalid friend user ID")
	}

	return engine.InsertOne(invitation)
}

//UpdateInvitation 更新邀请
func UpdateInvitation(invitation *entities.InvitationEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == invitation {
		return 0, errors.New("invalid parameter")
	}

	if invitation.ID <= 0 {
		return 0, errors.New("invalid invitation ID")
	}

	return engine.Id(invitation.ID).Update(invitation)
}

//GetInvitation 获取邀请
func GetInvitation(userID, friendUserID int64) (*entities.InvitationEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}

	if friendUserID <= 0 {
		return nil, errors.New("invalid friend user ID")
	}

	var invitation entities.InvitationEntity

	_, err := engine.Where("user_id = ?", userID).
		And("friend_user_id = ?", friendUserID).Get(&invitation)

	return &invitation, err
}

//GetInvitations 获取邀请列表
func GetInvitations(userID int64, status, pageIndex, pageSize int) (invitations []*entities.InvitationEntity, totalCount, pageCount int64, err error) {
	if nil == engine {
		err = errors.New("engine not initialized")
		return
	}

	if userID <= 0 {
		err = errors.New("invalid user ID")
		return
	}

	invitation := new(entities.InvitationEntity)
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
		rows, err = engine.Where("user_id = ?", userID).Limit(limit, offset).Rows(invitation)
	} else {
		rows, err = engine.Where("user_id = ?", userID).And("status = ?", status).Limit(limit, offset).Rows(invitation)
	}

	if err != nil {
		return
	}

	defer rows.Close()

	invitations = make([]*entities.InvitationEntity, 0)

	for rows.Next() {
		err = rows.Scan(invitation)
		if err != nil {
			return
		}

		invitations = append(invitations, invitation)

		//注意：这里应重新分配内存
		invitation = new(entities.InvitationEntity)
	}

	return
}

//UpdateInvitationStatus 更新邀请状态
func UpdateInvitationStatus(invitationID int64, status int) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if invitationID <= 0 {
		return 0, errors.New("invalid invitation ID")
	}

	return engine.Id(invitationID).Cols("status").Update(&entities.InvitationEntity{Status: status})
}
