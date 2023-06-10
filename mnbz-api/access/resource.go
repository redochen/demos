package access

import (
	"math"

	"github.com/redochen/demos/mnbz-api/models"
	"github.com/redochen/demos/mnbz-api/utils"

	"github.com/go-xorm/xorm"
	CcStr "github.com/redochen/tools/string"
)

// AddResource 添加资源
func AddResource(resource *models.Resource) (int32, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	if nil == resource {
		return 0, utils.NewInvalidError("resource")
	} else if resource.TripType == "" {
		return 0, utils.NewRequiredError("trip type")
	} else if resource.FromCity == "" {
		return 0, utils.NewRequiredError("from city")
	} else if resource.DepartDate == "" {
		return 0, utils.NewRequiredError("depart date")
	}

	if resource.GUID == "" {
		resource.GUID = utils.NewGUID()
	}

	_, err = engine.InsertOne(resource)
	if err != nil {
		return 0, err
	}

	found, err := getResourceByGUID(resource.GUID, false)
	if err != nil {
		return 0, err
	} else if nil == found || found.ID <= 0 {
		return 0, utils.NewFailedError("resource", "save")
	}

	return found.ID, nil
}

// UpdateResource 更新资源信息
func UpdateResource(model *models.Resource, userID int32) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	resource, err := getResource(model, true)
	if err != nil {
		return 0, err
	}

	//只允许操作自己的数据
	if userID != resource.UserID {
		return 0, utils.NewNotAllowedError("resource", "modify")
	}

	resource.TripType = CcStr.FirstValid(model.TripType, resource.TripType)
	resource.FromCity = CcStr.FirstValid(model.FromCity, resource.FromCity)
	resource.DepartDate = CcStr.FirstValid(model.DepartDate, resource.DepartDate)
	resource.BackDate = CcStr.FirstValid(model.BackDate, resource.BackDate)
	resource.Airlines = CcStr.FirstValid(model.Airlines, resource.Airlines)
	resource.Remarks = CcStr.FirstValid(model.Remarks, resource.Remarks)
	resource.Status = CcStr.FirstValid(model.Status, resource.Status)

	if model.AvailCount > 0 {
		resource.AvailCount = model.AvailCount
	}

	return engine.Id(resource.ID).Update(resource)
}

// DeleteResource 删除资源
func DeleteResource(model *models.Resource, userID int32) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	resource, err := getResource(model, true)
	if err != nil {
		return 0, err
	}

	//只允许操作自己的数据
	if userID != resource.UserID {
		return 0, utils.NewNotAllowedError("resource", "delete")
	}

	return engine.Id(resource.ID).Delete(resource)
}

// GetResource 获取资源
func GetResource(model *models.Resource, userID int32) (*models.Resource, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	resource, err := getResource(model, true)
	if err != nil {
		return nil, err
	}

	//只允许操作自己的数据
	if userID != resource.UserID {
		return nil, utils.NewNotAllowedError("resource", "query")
	}

	return resource, nil
}

// GetResources 获取资源列表
func GetResources(pageIndex, pageSize int, userID int32) (resources []*models.Resource, totalCount, pageCount int64, err error) {
	err = checkEngine()
	if err != nil {
		return
	}

	resource := new(models.Resource)
	offset, limit := utils.GetOffsetAndLimit(pageIndex, pageSize)

	totalCount, err = engine.Count(resource)
	if err != nil {
		return
	}

	if totalCount > 0 && limit > 0 {
		pageCount = int64(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var rows *xorm.Rows
	rows, err = engine.Where("userId = ?", userID).
		Desc("id").Limit(limit, offset).Rows(resource)
	if err != nil {
		return
	}

	defer rows.Close()

	resources = make([]*models.Resource, 0)

	for rows.Next() {
		err = rows.Scan(resource)
		if err != nil {
			return
		}

		resource.HandleResponse()
		resources = append(resources, resource)

		//注意：这里应重新分配内存
		resource = new(models.Resource)
	}

	return
}

// getResource 获取资源
func getResource(model *models.Resource, errIfNotExists bool) (*models.Resource, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if nil == model {
		return nil, utils.NewInvalidError("resource")
	}

	var resource *models.Resource

	if model.ID > 0 {
		resource, err = getResourceByID(model.ID, errIfNotExists)
	} else if model.GUID != "" {
		resource, err = getResourceByGUID(model.GUID, errIfNotExists)
	} else {
		err = utils.NewInvalidError("resource")
	}

	return resource, err
}

// getResourceByID 根据id获取资源
func getResourceByID(id int32, errIfNotExists bool) (*models.Resource, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if id <= 0 {
		return nil, utils.NewRequiredError("resource id")
	}

	resource := new(models.Resource)
	_, err = engine.Id(id).Get(resource)

	if err != nil {
		return nil, err
	} else if errIfNotExists && resource.ID <= 0 {
		return nil, utils.NewNotExistedError("resource")
	}

	resource.HandleResponse()

	return resource, err
}

// getResourceByGUID 根据guid获取资源
func getResourceByGUID(guid string, errIfNotExists bool) (*models.Resource, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if guid == "" {
		return nil, utils.NewRequiredError("resource guid")
	}

	resource := new(models.Resource)
	_, err = engine.Where("guid = ?", guid).Limit(1).Get(resource)

	if err != nil {
		return nil, err
	} else if errIfNotExists && resource.ID <= 0 {
		return nil, utils.NewNotExistedError("resource")
	}

	resource.HandleResponse()

	return resource, err
}
