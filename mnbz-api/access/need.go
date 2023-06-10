package access

import (
	"math"

	"github.com/redochen/demos/mnbz-api/models"
	"github.com/redochen/demos/mnbz-api/utils"

	"github.com/go-xorm/xorm"
	CcStr "github.com/redochen/tools/string"
)

// AddNeed 添加需求
func AddNeed(need *models.Need) (int32, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	if nil == need {
		return 0, utils.NewInvalidError("need")
	} else if need.TripType == "" {
		return 0, utils.NewRequiredError("trip type")
	} else if need.FromCity == "" {
		return 0, utils.NewRequiredError("from city")
	} else if need.DepartDate == "" {
		return 0, utils.NewRequiredError("depart date")
	}

	if need.GUID == "" {
		need.GUID = utils.NewGUID()
	}

	_, err = engine.InsertOne(need)
	if err != nil {
		return 0, err
	}

	found, err := getNeedByGUID(need.GUID, false)
	if err != nil {
		return 0, err
	} else if nil == found || found.ID <= 0 {
		return 0, utils.NewFailedError("need", "save")
	}

	return found.ID, nil
}

// UpdateNeed 更新需求信息
func UpdateNeed(model *models.Need, userID int32) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	need, err := getNeed(model, true)
	if err != nil {
		return 0, err
	}

	//只允许操作自己的数据
	if userID != need.UserID {
		return 0, utils.NewNotAllowedError("need", "modify")
	}

	need.TripType = CcStr.FirstValid(model.TripType, need.TripType)
	need.FromCity = CcStr.FirstValid(model.FromCity, need.FromCity)
	need.DepartDate = CcStr.FirstValid(model.DepartDate, need.DepartDate)
	need.BackDate = CcStr.FirstValid(model.BackDate, need.BackDate)
	need.Airlines = CcStr.FirstValid(model.Airlines, need.Airlines)
	need.Remarks = CcStr.FirstValid(model.Remarks, need.Remarks)
	need.Status = CcStr.FirstValid(model.Status, need.Status)

	if model.PassengerCount > 0 {
		need.PassengerCount = model.PassengerCount
	}

	return engine.Id(need.ID).Update(need)
}

// DeleteNeed 删除需求
func DeleteNeed(model *models.Need, userID int32) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	need, err := getNeed(model, true)
	if err != nil {
		return 0, err
	}

	//只允许操作自己的数据
	if userID != need.UserID {
		return 0, utils.NewNotAllowedError("need", "delete")
	}

	return engine.Id(need.ID).Delete(need)
}

// GetNeed 获取需求
func GetNeed(model *models.Need, userID int32) (*models.Need, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	need, err := getNeed(model, true)
	if err != nil {
		return nil, err
	}

	//只允许操作自己的数据
	if userID != need.UserID {
		return nil, utils.NewNotAllowedError("need", "query")
	}

	return need, nil
}

// GetNeeds 获取需求列表
func GetNeeds(pageIndex, pageSize int, userID int32) (needs []*models.Need, totalCount, pageCount int64, err error) {
	err = checkEngine()
	if err != nil {
		return
	}

	need := new(models.Need)
	offset, limit := utils.GetOffsetAndLimit(pageIndex, pageSize)

	totalCount, err = engine.Count(need)
	if err != nil {
		return
	}

	if totalCount > 0 && limit > 0 {
		pageCount = int64(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var rows *xorm.Rows
	rows, err = engine.Where("userId = ?", userID).
		Desc("id").Limit(limit, offset).Rows(need)
	if err != nil {
		return
	}

	defer rows.Close()

	needs = make([]*models.Need, 0)

	for rows.Next() {
		err = rows.Scan(need)
		if err != nil {
			return
		}

		need.HandleResponse()
		needs = append(needs, need)

		//注意：这里应重新分配内存
		need = new(models.Need)
	}

	return
}

// getNeed 获取需求
func getNeed(model *models.Need, errIfNotExists bool) (*models.Need, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if nil == model {
		return nil, utils.NewInvalidError("need")
	}

	var need *models.Need

	if model.ID > 0 {
		need, err = getNeedByID(model.ID, errIfNotExists)
	} else if model.GUID != "" {
		need, err = getNeedByGUID(model.GUID, errIfNotExists)
	} else {
		err = utils.NewInvalidError("need")
	}

	return need, err
}

// getNeedByID 根据id获取需求
func getNeedByID(id int32, errIfNotExists bool) (*models.Need, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if id <= 0 {
		return nil, utils.NewRequiredError("need id")
	}

	need := new(models.Need)
	_, err = engine.Id(id).Get(need)

	if err != nil {
		return nil, err
	} else if errIfNotExists && need.ID <= 0 {
		return nil, utils.NewNotExistedError("need")
	}

	need.HandleResponse()

	return need, err
}

// getNeedByGUID 根据guid获取需求
func getNeedByGUID(guid string, errIfNotExists bool) (*models.Need, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if guid == "" {
		return nil, utils.NewRequiredError("need guid")
	}

	need := new(models.Need)
	_, err = engine.Where("guid = ?", guid).Limit(1).Get(need)

	if err != nil {
		return nil, err
	} else if errIfNotExists && need.ID <= 0 {
		return nil, utils.NewNotExistedError("need")
	}

	need.HandleResponse()

	return need, err
}
