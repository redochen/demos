package config

import (
	"errors"
	. "github.com/redochen/tools/config"
	. "github.com/redochen/tools/function"
)

type PCC struct {
	UserId     string //账号
	Password   string //密码
	BranchCode string //Branch代码
	EndPoint   string //接口地址
}

type PccMap map[string]*PCC

var (
	PCCs PccMap //PCC列表
)

func InitPccs() {
	PCCs = make(map[string]*PCC)

	sections := Conf.GetSections()
	for _, section := range sections {
		if Conf.IsDefaultSection(section) {
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

func (m PccMap) SetPcc(pcc string) {
	if _, ok := m[pcc]; ok {
		panic("[" + pcc + "] already exists.")
	}

	userId, err := Conf.String(pcc, optionUserId)
	if err != nil {
		panic("[" + pcc + "] failed to load userId")
	}

	pwd, err := Conf.String(pcc, optionPassword)
	if err != nil {
		panic("[" + pcc + "] failed to load password")
	}

	branchCode, err := Conf.String(pcc, optionBranchCode)
	if err != nil {
		panic("[" + pcc + "] failed to load branchCode")
	}

	endPoint, err := Conf.String(pcc, optionEndPoint)
	if err != nil {
		panic("[" + pcc + "] failed to load endPoint")
	}

	ep, err := Endpoints.Get(endPoint)
	if err != nil {
		panic("[" + pcc + "] failed to get endpoint")
	}

	inst := &PCC{
		UserId:     userId,
		Password:   pwd,
		BranchCode: branchCode,
		EndPoint:   ep,
	}

	m[pcc] = inst
}

func (m PccMap) Get(pcc string) (*PCC, error) {
	defer CheckPanic()

	inst, ok := m[pcc]
	if !ok {
		return nil, errors.New("[" + pcc + "] not exists.")
	}

	return inst, nil
}
