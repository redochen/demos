package createpnr

import (
	"fmt"
	"time"

	CcTime "github.com/redochen/tools/time"
)

//Traveler 旅客
type Traveler struct {
	Surname       string     `xml:"surname" json:"surname"`                                 //姓氏
	GivenName     string     `xml:"givenName" json:"givenName"`                             //名字
	PassengerType string     `xml:"passengerType" json:"passengerType"`                     //类型：ADT-成人；CNN-儿童
	Gender        string     `xml:"gender" json:"gender"`                                   //性别：F-女；M-男
	Nationality   string     `xml:"nationality,omitempty" json:"nationality,omitempty"`     //国籍
	ReferenceName string     `xml:"referenceName,omitempty" json:"referenceName,omitempty"` //称呼
	BirthDate     string     `xml:"birthDate,omitempty" json:"birthDate,omitempty"`         //出生日期，格式为：YYYY-MM-DD
	Phones        []*Phone   `xml:"phones,omitempty" json:"phones,omitempty"`               //电话列表
	Emails        []*Email   `xml:"emails,omitempty" json:"emails,omitempty"`               //邮件列表
	Addresses     []*Address `xml:"addresses,omitempty" json:"addresses,omitempty"`         //地址列表
	FFPCarrier    string     `xml:"FFPCarrier,omitempty" json:"FFPCarrier,omitempty"`       //常旅客航司
	FFPCardNo     string     `xml:"FFPCardNo,omitempty" json:"FFPCardNo,omitempty"`         //常旅客卡号
	Key           string     `xml:"key,omitempty" json:"key,omitempty"`                     //引用键
	ElStat        string     `xml:"ElStat,omitempty" json:"ElStat,omitempty"`               //状态
}

//Phone 电话
type Phone struct {
	Location    string `xml:"location,omitempty" json:"location,omitempty"`       //所在地
	CountryCode string `xml:"countryCode,omitempty" json:"countryCode,omitempty"` //国家号
	AreaCode    string `xml:"areaCode,omitempty" json:"areaCode,omitempty"`       //区号
	Number      string `xml:"number" json:"number"`                               //电话
	Type        string `xml:"type,omitempty" json:"type,omitempty"`               //类型
}

//Email 邮件
type Email struct {
	Id   string `xml:"id" json:"id"`                         //邮件
	Type string `xml:"type,omitempty" json:"type,omitempty"` //类型
}

//Address 地址
type Address struct {
	Address    string   `xml:"address" json:"address"`                           //地址名称
	Country    string   `xml:"country,omitempty" json:"country,omitempty"`       //国家
	State      string   `xml:"state,omitempty" json:"state,omitempty"`           //州
	City       string   `xml:"city,omitempty" json:"city,omitempty"`             //市
	Street     []string `xml:"street,omitempty" json:"street,omitempty"`         //街道
	PostalCode string   `xml:"postalCode,omitempty" json:"postalCode,omitempty"` //邮编
	Key        string   `xml:"key,omitempty" json:"key,omitempty"`               //键值
	ElStat     string   `xml:"ElStat,omitempty" json:"ElStat,omitempty"`         //状态
}

//FormOfPayment 支付类型
type FormOfPayment struct {
	Type             string      `xml:"type,omitempty" json:"type,omitempty"`             //类型
	CreditCard       *CreditCard `xml:"creditCard,omitempty" json:"creditCard,omitempty"` //信用卡
	ProfileID        string      `xml:"profileID,omitempty" json:"profileID,omitempty"`
	ProfileKey       string      `xml:"profileKey,omitempty" json:"profileKey,omitempty"`
	ProviderInfoKeys []string    `xml:"providerInfoKeys,omitempty" json:"providerInfoKeys,omitempty"` //供应商预订引用键
	Reusable         bool        `xml:"reusable,omitempty" json:"reusable,omitempty"`                 //是否可重复使用
	Key              string      `xml:"key,omitempty" json:"key,omitempty"`                           //键值
	ElStat           string      `xml:"ElStat,omitempty" json:"ElStat,omitempty"`                     //状态
}

//CreditCard 信用卡参数类
type CreditCard struct {
	Type           string   `xml:"type,omitempty" json:"type,omitempty"`                     //类型
	CountryCode    string   `xml:"countryCode" json:"countryCode"`                           //国家代码
	StateCode      string   `xml:"stateCode" json:"stateCode"`                               //状态代码
	BankName       string   `xml:"bankName" json:"bankName"`                                 //银行名称
	Number         string   `xml:"number" json:"number"`                                     //卡号
	ExpiryDate     string   `xml:"expiryDate,omitempty" json:"expiryDate,omitempty"`         //过期日期
	CVV            string   `xml:"CVV,omitempty" json:"CVV,omitempty"`                       //CVV号
	ApprovalCode   string   `xml:"approvalCode,omitempty" json:"approvalCode,omitempty"`     //批准代码
	BillingAddress *Address `xml:"billingAddress,omitempty" json:"billingAddress,omitempty"` //账单地址
	Key            string   `xml:"key,omitempty" json:"key,omitempty"`                       //键值
	ProfileID      string   `xml:"profileID,omitempty" json:"profileID,omitempty"`
}

//OtherServiceInfo 其他服务信息
type OtherServiceInfo struct {
	Carrier         string `xml:"carrier,omitempty" json:"carrier,omitempty"`                 //？？
	Code            string `xml:"code,omitempty" json:"code,omitempty"`                       //代码
	Text            string `xml:"text" json:"text"`                                           //文本
	ProviderInfoRef string `xml:"providerInfoRef,omitempty" json:"providerInfoRef,omitempty"` //供应商预订引用键
	ProviderCode    string `xml:"providerCode,omitempty" json:"providerCode,omitempty"`       //供应商代码
	Key             string `xml:"key,omitempty" json:"key,omitempty"`                         //键值
	ElStat          string `xml:"ElStat,omitempty" json:"ElStat,omitempty"`                   //状态
}

//HotelProperty 酒店物业
type HotelProperty struct {
	Chain              string   `xml:"chain" json:"chain"`                                           //酒店连锁代码
	Code               string   `xml:"code" json:"code"`                                             //酒店代码
	Location           string   `xml:"location" json:"location"`                                     //酒店所在地
	Name               string   `xml:"name" json:"name"`                                             //酒店名称
	ParticipationLevel string   `xml:"participantLevel,omitempty" json:"participantLevel,omitempty"` //参与级别
	Addresses          []string `xml:"addresses,omitempty" json:"addresses,omitempty"`               //酒店地址列表
	Phones             []*Phone `xml:"phones,omitempty" json:"phones,omitempty"`                     //电话列表
	Availability       string   `xml:"availability,omitempty" json:"availability,omitempty"`         //可用性
}

//HotelRate 酒店价格
type HotelRate struct {
	Base          float32            `xml:"base" json:"base"`                                       //基本价格
	BaseCurrency  string             `xml:"baseCurrency" json:"baseCurrency"`                       //基本价格货币代码
	Tax           float32            `xml:"tax,omitempty" json:"tax,omitempty"`                     //税费
	TaxCurrency   string             `xml:"taxCurrency,omitempty" json:"taxCurrency,omitempty"`     //税费货币代码
	Total         float32            `xml:"total,omitempty" json:"total,omitempty"`                 //总价
	TotalCurrency string             `xml:"totalCurrency,omitempty" json:"totalCurrency,omitempty"` //总价货币代码
	RatePlanType  string             `xml:"ratePlanType,omitempty" json:"ratePlanType,omitempty"`   //价格编码
	Category      int                `xml:"category,omitempty" json:"category,omitempty"`           //价格类别
	Guaranteed    bool               `xml:"guaranteed,omitempty" json:"guaranteed,omitempty"`       //是否保证价格
	Descriptions  []string           `xml:"descriptions,omitempty" json:"descriptions,omitempty"`   //价格描述
	RateByDates   []*HotelRateByDate `xml:"rateByDates,omitempty" json:"rateByDates,omitempty"`     //限期价格列表
}

//HotelRateByDate 限期价格
type HotelRateByDate struct {
	Base          float32 `xml:"base" json:"base"`                                       //基本价格
	BaseCurrency  string  `xml:"baseCurrency" json:"baseCurrency"`                       //基本价格货币代码
	Tax           float32 `xml:"tax,omitempty" json:"tax,omitempty"`                     //税费
	TaxCurrency   string  `xml:"taxCurrency,omitempty" json:"taxCurrency,omitempty"`     //税费货币代码
	Total         float32 `xml:"total,omitempty" json:"total,omitempty"`                 //总价
	TotalCurrency string  `xml:"totalCurrency,omitempty" json:"totalCurrency,omitempty"` //总价货币代码
	EffectiveDate string  `xml:"effectiveDate,omitempty" json:"effectiveDate,omitempty"` //生效日期
	ExpireDate    string  `xml:"expireDate,omitempty" json:"expireDate,omitempty"`       //过期日期
	Contents      string  `xml:"contents,omitempty" json:"contents,omitempty"`           //内容
}

//GetChildYear v获取儿童年龄
func (traveler *Traveler) GetChildYear(isRemark bool) string {
	if len(traveler.BirthDate) == 0 {
		return ""
	}

	//YYYY-MM-DD
	format := "2006-01-02"
	dt, err := time.Parse(format, traveler.BirthDate)
	if err != nil {
		return ""
	}

	diff := CcTime.CalculateDateDifference(dt, time.Now())
	if nil == diff {
		return ""
	}

	if isRemark {
		return fmt.Sprintf("P-C%02d", diff.Years)
	}

	return fmt.Sprintf("C%02d", diff.Years)
}
