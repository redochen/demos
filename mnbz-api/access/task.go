package access

import (
	"math"

	"github.com/redochen/demos/mnbz-api/models"
	"github.com/redochen/demos/mnbz-api/utils"

	"github.com/go-xorm/xorm"
	CcStr "github.com/redochen/tools/string"
)

// AddTask 添加抢票
func AddTask(task *models.Task) (int32, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	if nil == task {
		return 0, utils.NewInvalidError("task")
	} else if task.AvCmd == "" {
		return 0, utils.NewRequiredError("av cmd")
	}
	//  else if task.BookCmd == "" {
	// 	return 0, utils.NewRequiredError("book cmd")
	// } else if task.PassengerNames == "" {
	// 	return 0, utils.NewRequiredError("passenger names")
	// } else if task.PassengerDocs == "" {
	// 	return 0, utils.NewRequiredError("passenger docs")
	// }

	if task.GUID == "" {
		task.GUID = utils.NewGUID()
	}

	_, err = engine.InsertOne(task)
	if err != nil {
		return 0, err
	}

	found, err := getTaskByGUID(task.GUID, false)
	if err != nil {
		return 0, err
	} else if nil == found || found.ID <= 0 {
		return 0, utils.NewFailedError("task", "save")
	}

	return found.ID, nil
}

// UpdateTask 更新抢票信息
func UpdateTask(model *models.Task, userID int32) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	task, err := getTask(model, true)
	if err != nil {
		return 0, err
	}

	//只允许操作自己的数据
	if userID != task.UserID {
		return 0, utils.NewNotAllowedError("task", "modify")
	}

	task.PNR = CcStr.FirstValid(model.PNR, task.PNR)
	task.AvCmd = CcStr.FirstValid(model.AvCmd, task.AvCmd)
	task.BookCmd = CcStr.FirstValid(model.BookCmd, task.BookCmd)
	task.PassengerNames = CcStr.FirstValid(model.PassengerNames, task.PassengerNames)
	task.PassengerDocs = CcStr.FirstValid(model.PassengerDocs, task.PassengerDocs)
	task.Status = CcStr.FirstValid(model.Status, task.Status)

	if model.MaxCount > 0 {
		task.MaxCount = model.MaxCount
	}

	if model.Interval > 0 {
		task.Interval = model.Interval
	}

	return engine.Id(task.ID).Update(task)
}

// DeleteTask 删除抢票
func DeleteTask(model *models.Task, userID int32) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	task, err := getTask(model, true)
	if err != nil {
		return 0, err
	}

	//只允许操作自己的数据
	if userID != task.UserID {
		return 0, utils.NewNotAllowedError("task", "delete")
	}

	return engine.Id(task.ID).Delete(task)
}

// GetTask 获取抢票
func GetTask(model *models.Task, userID int32) (*models.Task, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	task, err := getTask(model, true)
	if err != nil {
		return nil, err
	}

	//只允许操作自己的数据
	if userID != task.UserID {
		return nil, utils.NewNotAllowedError("task", "query")
	}

	return task, nil
}

// GetTasks 获取抢票列表
func GetTasks(pageIndex, pageSize int, userID int32) (tasks []*models.Task, totalCount, pageCount int64, err error) {
	err = checkEngine()
	if err != nil {
		return
	}

	task := new(models.Task)
	offset, limit := utils.GetOffsetAndLimit(pageIndex, pageSize)

	totalCount, err = engine.Count(task)
	if err != nil {
		return
	}

	if totalCount > 0 && limit > 0 {
		pageCount = int64(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var rows *xorm.Rows
	rows, err = engine.Where("userId = ?", userID).
		Desc("id").Limit(limit, offset).Rows(task)
	if err != nil {
		return
	}

	defer rows.Close()

	tasks = make([]*models.Task, 0)

	for rows.Next() {
		err = rows.Scan(task)
		if err != nil {
			return
		}

		task.HandleResponse()
		tasks = append(tasks, task)

		//注意：这里应重新分配内存
		task = new(models.Task)
	}

	return
}

// getTask 获取抢票
func getTask(model *models.Task, errIfNotExists bool) (*models.Task, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if nil == model {
		return nil, utils.NewInvalidError("task")
	}

	var task *models.Task

	if model.ID > 0 {
		task, err = getTaskByID(model.ID, errIfNotExists)
	} else if model.GUID != "" {
		task, err = getTaskByGUID(model.GUID, errIfNotExists)
	} else {
		err = utils.NewInvalidError("task")
	}

	return task, err
}

// getTaskByID 根据id获取抢票
func getTaskByID(id int32, errIfNotExists bool) (*models.Task, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if id <= 0 {
		return nil, utils.NewRequiredError("task id")
	}

	task := new(models.Task)
	_, err = engine.Id(id).Get(task)

	if err != nil {
		return nil, err
	} else if errIfNotExists && task.ID <= 0 {
		return nil, utils.NewNotExistedError("task")
	}

	task.HandleResponse()

	return task, err
}

// getTaskByGUID 根据guid获取抢票
func getTaskByGUID(guid string, errIfNotExists bool) (*models.Task, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if guid == "" {
		return nil, utils.NewRequiredError("task guid")
	}

	task := new(models.Task)
	_, err = engine.Where("guid = ?", guid).Limit(1).Get(task)

	if err != nil {
		return nil, err
	} else if errIfNotExists && task.ID <= 0 {
		return nil, utils.NewNotExistedError("task")
	}

	task.HandleResponse()

	return task, err
}
