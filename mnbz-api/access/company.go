package access

import (
	"math"

	"github.com/redochen/demos/mnbz-api/models"
	"github.com/redochen/demos/mnbz-api/utils"

	"github.com/go-xorm/xorm"
	CcStr "github.com/redochen/tools/string"
)

// AddCompany 添加公司
func AddCompany(company *models.Company, checkExist bool) (int32, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	if nil == company {
		return 0, utils.NewInvalidError("company")
	} else if company.GUID == "" {
		company.GUID = utils.NewGUID()
		//return 0, utils.NewRequiredError("company guid")
	}
	//  else if company.Country == "" {
	// 	return 0, utils.NewRequiredError("company country")
	// } else if company.City == "" {
	// 	return 0, utils.NewRequiredError("company city")
	// }

	found, err := GetCompanyByName(company.Name, false)
	if err != nil {
		return 0, err
	} else if found != nil && found.ID > 0 {
		if checkExist {
			return 0, utils.NewExistedError("company")
		} else {
			return found.ID, nil
		}
	}

	_, err = engine.InsertOne(company)
	if err != nil {
		return 0, err
	}

	found, err = GetCompanyByGUID(company.GUID, false)
	if err != nil {
		return 0, err
	} else if nil == found || found.ID <= 0 {
		return 0, utils.NewFailedError("company", "save")
	}

	return found.ID, nil
}

// UpdateCompany 更新公司信息
func UpdateCompany(company *models.Company) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	if nil == company {
		return 0, utils.NewInvalidError("company")
	}

	found, err := GetCompanyByGUID(company.GUID, true)
	if err != nil {
		return 0, err
	}

	company.Name = CcStr.FirstValid(company.Name, found.Name)
	company.Country = CcStr.FirstValid(company.Country, found.Country)
	company.City = CcStr.FirstValid(company.City, found.City)
	company.ContractUser = CcStr.FirstValid(company.ContractUser, found.ContractUser)

	return engine.Id(company.ID).Update(company)
}

// DeleteCompany 删除公司
func DeleteCompany(guid string) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	company, err := GetCompanyByGUID(guid, true)
	if err != nil {
		return 0, err
	}

	model := new(models.Company)
	return engine.Id(company.ID).Delete(model)
}

// GetCompanys 获取公司列表
func GetCompanys(pageIndex, pageSize int) (companys []*models.Company, totalCount, pageCount int64, err error) {
	err = checkEngine()
	if err != nil {
		return
	}

	company := new(models.Company)
	offset, limit := utils.GetOffsetAndLimit(pageIndex, pageSize)

	totalCount, err = engine.Count(company)
	if err != nil {
		return
	}

	if totalCount > 0 && limit > 0 {
		pageCount = int64(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var rows *xorm.Rows
	rows, err = engine.Desc("id").Limit(limit, offset).Rows(company)
	if err != nil {
		return
	}

	defer rows.Close()

	companys = make([]*models.Company, 0)

	for rows.Next() {
		err = rows.Scan(company)
		if err != nil {
			return
		}

		company.HandleResponse()
		companys = append(companys, company)

		//注意：这里应重新分配内存
		company = new(models.Company)
	}

	return
}

// GetCompanyByGUID 根据guid获取公司
func GetCompanyByGUID(guid string, errIfNotExists bool) (*models.Company, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if guid == "" {
		return nil, utils.NewRequiredError("company guid")
	}

	company := new(models.Company)
	_, err = engine.Where("guid = ?", guid).Limit(1).Get(company)

	if err != nil {
		return nil, err
	} else if errIfNotExists && company.ID <= 0 {
		return nil, utils.NewNotExistedError("company")
	}

	company.HandleResponse()

	return company, err
}

// GetCompanyByName 根据公司名称获取公司信息
func GetCompanyByName(companyName string, errIfNotExists bool) (*models.Company, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if companyName == "" {
		return nil, utils.NewRequiredError("company name")
	}

	company := new(models.Company)
	_, err = engine.Where("name = ?", companyName).Limit(1).Get(company)

	if err != nil {
		return nil, err
	} else if errIfNotExists && company.ID <= 0 {
		return nil, utils.NewNotExistedError("company")
	}

	company.HandleResponse()

	return company, err
}
