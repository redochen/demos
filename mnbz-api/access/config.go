package access

import (
	"math"

	"github.com/redochen/demos/mnbz-api/models"
	"github.com/redochen/demos/mnbz-api/utils"

	"github.com/go-xorm/xorm"
	CcStr "github.com/redochen/tools/string"
)

// AddConfig 添加配置
func AddConfig(config *models.Config) (int32, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	if nil == config {
		return 0, utils.NewInvalidError("config")
	}
	// else if config.Account == "" {
	// 	return 0, utils.NewRequiredError("config account")
	// } else if config.Password == "" {
	// 	return 0, utils.NewRequiredError("config password")
	// }

	if config.GUID == "" {
		config.GUID = utils.NewGUID()
	}

	found, err := getConfigByGdsTypeAndPCC(config.GdsType, config.PCC, false)
	if err != nil {
		return 0, err
	} else if found != nil && found.ID > 0 {
		return 0, utils.NewExistedError("config")
	}

	_, err = engine.InsertOne(config)
	if err != nil {
		return 0, err
	}

	found, err = getConfigByGUID(config.GUID, false)
	if err != nil {
		return 0, err
	} else if nil == found || found.ID <= 0 {
		return 0, utils.NewFailedError("config", "save")
	}

	return found.ID, nil
}

// UpdateConfig 更新配置信息
func UpdateConfig(model *models.Config, userID int32) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	config, err := getConfig(model, true)
	if err != nil {
		return 0, err
	}

	//只允许操作自己的数据
	if userID != config.UserID {
		return 0, utils.NewNotAllowedError("config", "modify")
	}

	config.GdsType = CcStr.FirstValid(model.GdsType, config.GdsType)
	config.PCC = CcStr.FirstValid(model.PCC, config.PCC)
	config.Account = CcStr.FirstValid(model.Account, config.Account)
	config.Password = CcStr.FirstValid(model.Password, config.Password)
	config.Status = CcStr.FirstValid(model.Status, config.Status)

	return engine.Id(config.ID).Update(config)
}

// DeleteConfig 删除配置
func DeleteConfig(model *models.Config, userID int32) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	config, err := getConfig(model, true)
	if err != nil {
		return 0, err
	}

	//只允许操作自己的数据
	if userID != config.UserID {
		return 0, utils.NewNotAllowedError("config", "delete")
	}

	return engine.Id(config.ID).Delete(config)
}

// GetConfig 获取配置
func GetConfig(model *models.Config, userID int32) (*models.Config, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	config, err := getConfig(model, true)
	if err != nil {
		return nil, err
	}

	//只允许操作自己的数据
	if userID != config.UserID {
		return nil, utils.NewNotAllowedError("config", "query")
	}

	return config, nil
}

// GetConfigs 获取配置列表
func GetConfigs(pageIndex, pageSize int, userID int32) (configs []*models.Config, totalCount, pageCount int64, err error) {
	err = checkEngine()
	if err != nil {
		return
	}

	config := new(models.Config)
	offset, limit := utils.GetOffsetAndLimit(pageIndex, pageSize)

	totalCount, err = engine.Count(config)
	if err != nil {
		return
	}

	if totalCount > 0 && limit > 0 {
		pageCount = int64(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var rows *xorm.Rows
	rows, err = engine.Where("userId = ?", userID).
		Desc("id").Limit(limit, offset).Rows(config)
	if err != nil {
		return
	}

	defer rows.Close()

	configs = make([]*models.Config, 0)

	for rows.Next() {
		err = rows.Scan(config)
		if err != nil {
			return
		}

		config.HandleResponse()
		configs = append(configs, config)

		//注意：这里应重新分配内存
		config = new(models.Config)
	}

	return
}

// getConfig 获取配置
func getConfig(model *models.Config, errIfNotExists bool) (*models.Config, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if nil == model {
		return nil, utils.NewInvalidError("config")
	}

	var config *models.Config

	if model.ID > 0 {
		config, err = getConfigByID(model.ID, errIfNotExists)
	} else if model.GUID != "" {
		config, err = getConfigByGUID(model.GUID, errIfNotExists)
	} else if model.GdsType != "" && model.PCC != "" {
		config, err = getConfigByGdsTypeAndPCC(model.GdsType, model.PCC, errIfNotExists)
	} else {
		err = utils.NewInvalidError("config")
	}

	return config, err
}

// getConfigByID 根据id获取配置
func getConfigByID(id int32, errIfNotExists bool) (*models.Config, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if id <= 0 {
		return nil, utils.NewRequiredError("config id")
	}

	config := new(models.Config)
	_, err = engine.Id(id).Get(config)

	if err != nil {
		return nil, err
	} else if errIfNotExists && config.ID <= 0 {
		return nil, utils.NewNotExistedError("config")
	}

	config.HandleResponse()

	return config, err
}

// getConfigByGUID 根据guid获取配置
func getConfigByGUID(guid string, errIfNotExists bool) (*models.Config, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if guid == "" {
		return nil, utils.NewRequiredError("config guid")
	}

	config := new(models.Config)
	_, err = engine.Where("guid = ?", guid).Limit(1).Get(config)

	if err != nil {
		return nil, err
	} else if errIfNotExists && config.ID <= 0 {
		return nil, utils.NewNotExistedError("config")
	}

	config.HandleResponse()

	return config, err
}

// getConfigByGdsTypeAndPCC 根据GDS类型和PCC获取配置
func getConfigByGdsTypeAndPCC(gdsType, pcc string, errIfNotExists bool) (*models.Config, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if gdsType == "" {
		return nil, utils.NewRequiredError("gds type")
	}

	if pcc == "" {
		return nil, utils.NewRequiredError("pcc")
	}

	config := new(models.Config)
	_, err = engine.Where("gdsType = ?", gdsType).And("pcc = ?", pcc).Limit(1).Get(config)

	if err != nil {
		return nil, err
	} else if errIfNotExists && config.ID <= 0 {
		return nil, utils.NewNotExistedError("config")
	}

	config.HandleResponse()

	return config, err
}
