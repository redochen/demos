package common

import (
	CcStr "github.com/redochen/tools/string"
)

//GetCurrency 获取货币代码
func (typeMoney TypeMoney) GetCurrency() string {
	money := string(typeMoney)
	if len(money) > 3 {
		return money[0:3]
	}
	return ""
}

//GetAmount 获取金额
func (typeMoney TypeMoney) GetAmount() int {
	money := string(typeMoney)
	if len(money) > 3 {
		return int(CcStr.ParseInt(money[3:]))
	}
	return int(CcStr.ParseInt(money))
}

//GetFloatAmount 获取金额
func (typeMoney TypeMoney) GetFloatAmount() float32 {
	money := string(typeMoney)
	if len(money) > 3 {
		return float32(CcStr.ParseFloat(money[3:]))
	}
	return float32(CcStr.ParseFloat(money))
}
