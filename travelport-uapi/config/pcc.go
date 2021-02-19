package config

import (
	"errors"

	"github.com/redochen/tools/config"
	CcFunc "github.com/redochen/tools/function"
)

//PCC pcc上下文
type PCC struct {
	UserID     string //账号
	Password   string //密码
	BranchCode string //Branch代码
	EndPoint   string //接口地址
}

//PccMap pcc集合定义
type PccMap map[string]*PCC

var (
	//PCCs PCC列表
	PCCs PccMap
)

//InitPccs 初始化pcc列表
func InitPccs() {
	PCCs = make(map[string]*PCC)

	sections := config.Conf.GetSections()
	for _, section := range sections {
		if config.Conf.IsDefaultSection(section) {
			continue
		}

		if section == sectionSvr ||
			section == sectionLog ||
			section == sectionEndpoints {
			continue
		}

		PCCs.SetPcc(section)
	}

	//fmt.Printf("%v\n", PCCs)
}

//SetPcc 设置pcc配置
func (m PccMap) SetPcc(pcc string) {
	if _, ok := m[pcc]; ok {
		panic("[" + pcc + "] already exists.")
	}

	userID, err := config.Conf.String(pcc, optionUserID)
	if err != nil {
		panic("[" + pcc + "] failed to load userId")
	}

	pwd, err := config.Conf.String(pcc, optionPassword)
	if err != nil {
		panic("[" + pcc + "] failed to load password")
	}

	branchCode, err := config.Conf.String(pcc, optionBranchCode)
	if err != nil {
		panic("[" + pcc + "] failed to load branchCode")
	}

	endPoint, err := config.Conf.String(pcc, optionEndPoint)
	if err != nil {
		panic("[" + pcc + "] failed to load endPoint")
	}

	ep, err := Endpoints.Get(endPoint)
	if err != nil {
		panic("[" + pcc + "] failed to get endpoint")
	}

	inst := &PCC{
		UserID:     userID,
		Password:   pwd,
		BranchCode: branchCode,
		EndPoint:   ep,
	}

	m[pcc] = inst
}

//Get 根据pcc获取配置
func (m PccMap) Get(pcc string) (*PCC, error) {
	defer CcFunc.CheckPanic()

	inst, ok := m[pcc]
	if !ok {
		return nil, errors.New("[" + pcc + "] not exists.")
	}

	return inst, nil
}
