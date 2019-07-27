package hotel

//酒店详情项目
type HotelDetailItem struct {
	Name string   `xml:"Name" json:"name"` //名称
	Text []string `xml:"Text" json:"text"` //内容
}

//酒店规则项目
type HotelRuleItem struct {
	Name string   `xml:"Name" json:"name"` //名称
	Text []string `xml:"Text" json:"text"` //内容
}

//酒店属性
type HotelProperty struct {
	Chain                 string        `xml:"chain" json:"chain"`                                 //所属连锁酒店名称
	Code                  string        `xml:"code" json:"code"`                                   //酒店代码
	Name                  string        `xml:"name" json:"name"`                                   //酒店名称
	Location              string        `xml:"location" json:"location"`                           //酒店地址
	Address               []string      `xml:"address" json:"address"`                             //物业地址
	Phones                []*PhoneInfo  `xml:"phones" json:"phones"`                               //电话
	Latitude              float64       `xml:"latitude" json:"latitude"`                           //纬度
	Longitude             float64       `xml:"longitude" json:"longitude"`                         // 经度
	DistanceUnits         string        `xml:"distanceUnits" json:"distanceUnits"`                 //距离单位
	DistanceValue         int           `xml:"distanceValue" json:"distanceValue"`                 //距离值
	DistanceDirection     string        `xml:"distanceDirection" json:"distanceDirection"`         //距离方向
	Ratings               []*RatingInfo `xml:"ratings" json:"ratings"`                             //评级列表
	Amenities             []string      `xml:"amenities" json:"amenities"`                         //设施列表
	Transportation        string        `xml:"transportation" json:"transportation"`               //交通设施
	ReserveRequirement    string        `xml:"reserveRequirement" json:"reserveRequirement"`       //预订要求
	Availability          string        `xml:"availability" json:"availability"`                   //可用性
	PreferredOption       bool          `xml:"preferredOption" json:"preferredOption"`             //是否首选
	FeaturedProperty      bool          `xml:"featuredProperty" json:"featuredProperty"`           //是否为特性物业？
	ParticipationLevel    string        `xml:"participationLevel" json:"participationLevel"`       //参与级别？
	NetTransCommissionInd string        `xml:"netTransCommissionInd" json:"netTransCommissionInd"` //
	Key                   string        `xml:"key" json:"key"`                                     //键值
	VendorLocationKey     string        `xml:"vendorLocationKey" json:"vendorLocationKey"`         //供应商地址键值
}

//酒店价格详情
type HotelRateDetail struct {
	Descriptions              []*RoomRateDescription `xml:"descriptions" json:"descriptions"`                           //描述列表
	DateRates                 []*HotelRateByDate     `xml:"dateRates" json:"dateRates"`                                 //日期价格列表
	MatchInds                 []*RateMatchIndicator  `xml:"matchInds" json:"matchInds"`                                 //价格匹配指标器列表
	TaxDetails                []*TaxDetail           `xml:"taxDetails" json:"taxDetails"`                               //税费详情
	MealPlans                 []string               `xml:"mealPlans" json:"mealPlans"`                                 //餐食列表
	CancelRule                *CancelRule            `xml:"cancelRule" json:"cancelRule"`                               //退订规定
	GuaranteeInfo             *GuaranteeInfo         `xml:"guaranteeInfo" json:"guaranteeInfo"`                         //担保信息
	RoomView                  string                 `xml:"roomView" json:"roomView"`                                   //房间景观
	SupplementalRateInfo      string                 `xml:"supplementalRateInfo" json:"supplementalRateInfo"`           //供应商价格信息
	RatePlanType              string                 `xml:"ratePlanType" json:"ratePlanType"`                           //价格计划类型
	Rate                      *RateInfo              `xml:"rate" json:"rate"`                                           //价格信息
	RateGuaranteed            bool                   `xml:"rateGuaranteed" json:"rateGuaranteed"`                       //价格是否已担保
	ApproximateRate           *RateInfo              `xml:"approximateRate" json:"approximateRate"`                     //近似价格信息
	ApproximateRateGuaranteed bool                   `xml:"approximateRateGuaranteed" json:"approximateRateGuaranteed"` //近似价格是否已担保
	RateCategory              string                 `xml:"rateCategory" json:"rateCategory"`                           //价格分类
	Key                       string                 `xml:"key" json:"key"`                                             //
	RateSupplier              string                 `xml:"rateSupplier" json:"rateSupplier"`                           //价格供应商
	RateOfferId               string                 `xml:"rateOfferId" json:"rateOfferId"`                             //
	InPolicy                  bool                   `xml:"inPolicy" json:"inPolicy"`                                   //
}

type RoomRateDescription struct {
	Name string   `xml:"name" json:"name"` //名称
	Text []string `xml:"text" json:"text"` //内容
}

//日期价格
type HotelRateByDate struct {
	EffectiveDate   string    `xml:"effectiveDate" json:"effectiveDate"`     //生效日期
	ExpireDate      string    `xml:"expireDate" json:"expireDate"`           //过期日期
	Rate            *RateInfo `xml:"rate" json:"rate"`                       //价格信息
	ApproximateRate *RateInfo `xml:"approximateRate" json:"approximateRate"` //近似价格信息
	Contents        string    `xml:"contents" json:"contents"`               //其他内容
}

//评级信息
type RatingInfo struct {
	Type    string `xml:"type" json:"type"`       //类型
	Ratings []int  `xml:"ratings" json:"ratings"` //评级列表
}

//价格信息
type RateInfo struct {
	Base              float32 `xml:"base" json:"base"`                           //基础价格
	BaseCurrency      string  `xml:"baseCurrency" json:"baseCurrency"`           //基础价格货币代码
	Tax               float32 `xml:"tax" json:"tax"`                             //税费
	TaxCurrency       string  `xml:"taxCurrency" json:"taxCurrency"`             //税费货币代码
	Total             float32 `xml:"total" json:"total"`                         //总价格
	TotalCurrency     string  `xml:"totalCurrency" json:"totalCurrency"`         //总价格货币代码
	Surcharge         float32 `xml:"surcharge" json:"surcharge"`                 //附加费用
	SurchargeCurrency string  `xml:"surchargeCurrency" json:"surchargeCurrency"` //附加费用货币代码
}

//电话信息
type PhoneInfo struct {
	Key         string `xml:"key" json:"key"`                 //键值
	Type        string `xml:"type" json:"type"`               //类型
	Location    string `xml:"location" json:"location"`       //所在城市
	CountryCode string `xml:"countryCode" json:"countryCode"` //国家代码
	AreaCode    string `xml:"areaCode" json:"areaCode"`       //地区代码
	Number      string `xml:"number" json:"number"`           //号码
	Extension   string `xml:"extension" json:"extension"`     //扩展信息
	Text        string `xml:"text" json:"text"`               //相关信息
}

//价格匹配指标器
type RateMatchIndicator struct {
	Type   string `xml:"type" json:"type"`
	Status string `xml:"status" json:"status"`
	Value  string `xml:"value" json:"value"`
}

//税费详情
type TaxDetail struct {
	Amount         float32 `xml:"amount" json:"amount"`                 //金额
	Currency       string  `xml:"currency" json:"currency"`             //货币代码
	Percentage     float32 `xml:"percentage" json:"percentage"`         //百分比
	Code           int     `xml:"code" json:"code"`                     //代码
	EffectiveDate  string  `xml:"effectiveDate" json:"effectiveDate"`   //生效日期
	ExpireDate     string  `xml:"expireDate" json:"expireDate"`         //过期日期
	Term           string  `xml:"term" json:"term"`                     //收取方式：PerPerson, PerNight and PerStay
	CollectionFreq string  `xml:"collectionFreq" json:"collectionFreq"` //收取频率：Once or Daily
}

//退订规则
type CancelRule struct {
	NonRefundable                 string  `xml:"nonRefundable" json:"nonRefundable"`                                 //是否不可退订
	CancelDeadline                string  `xml:"cancelDeadline" json:"cancelDeadline"`                               //免费退订最后期限
	TaxInclusive                  bool    `xml:"taxInclusive" json:"taxInclusive"`                                   //退订费中是否包含税费
	FeeInclusive                  bool    `xml:"feeInclusive" json:"feeInclusive"`                                   //退订费中是否包含手续费
	CancelPenaltyAmount           float32 `xml:"cancelPenaltyAmount" json:"cancelPenaltyAmount"`                     //退订费用金额
	CancelPenaltyAmountCurrency   string  `xml:"cancelPenaltyAmountCurrency" json:"cancelPenaltyAmountCurrency"`     //退订费用货币代码
	NumberOfNights                uint    `xml:"numberOfNights" json:"numberOfNights"`                               //几夜？？
	CancelPenaltyPercent          float32 `xml:"cancelPenaltyPercent" json:"cancelPenaltyPercent"`                   //退订费用百分比
	CancelPenaltyPercentAppliesTo string  `xml:"cancelPenaltyPercentAppliesTo" json:"cancelPenaltyPercentAppliesTo"` //退订费用百分比限定符？？
}

//担保信息
type GuaranteeInfo struct {
	DepositAmount                    float32                 `xml:"depositAmount" json:"depositAmount"`                                       //预付金额
	DepositAmountCurrency            string                  `xml:"depositAmountCurrency" json:"depositAmountCurrency"`                       //预付金额货币代码
	ApproximateDepositAmount         float32                 `xml:"approximateDepositAmount" json:"approximateDepositAmount"`                 //近似预付金额
	ApproximateDepositAmountCurrency string                  `xml:"approximateDepositAmountCurrency" json:"approximateDepositAmountCurrency"` //近似预付金额货币代码
	DepositNights                    int                     `xml:"depositNights" json:"dpositNights"`                                        //预付几夜？？
	DepositPercent                   int                     `xml:"depositPercent" json:"depositPercent"`                                     //预付百分比
	GuaranteePaymentType             []*GuaranteePaymentType `xml:"guaranteePaymentType" json:"guaranteePaymentType"`                         //担保支付类型列表
	AbsoluteDeadline                 string                  `xml:"absoluteDeadline" json:"absoluteDeadline"`                                 //绝对最后期限
	CredentialsRequired              bool                    `xml:"credentialsRequired" json:"credentialsRequired"`                           //预订时是否需要凭证
	HoldTime                         string                  `xml:"holdTime" json:"holdTime"`                                                 //房间保留时间
	GuaranteeType                    string                  `xml:"guaranteeType" json:"guaranteeType"`                                       //担保类型：Deposit, Guarantee, or Prepayment
}

//担保支付类型
type GuaranteePaymentType struct {
	Value       string `xml:"value" json:"value"`             //担保支付
	Type        string `xml:"type" json:"type"`               //担保支付类型：CreditCard, AgencyIATA/ARC, FrequentGuest, SpecialIndustry, CDNumber, HomeAddress, CompanyAddress, Override, Other, or None
	Description string `xml:"description" json:"description"` //描述
}

//客户评论
type Comments struct {
	Content   string `xml:"content" json:"content"`     //评论内容
	Date      string `xml:"date" json:"date"`           //评论时间
	Language  string `xml:"language" json:"language"`   //语言
	Source    string `xml:"source" json:"source"`       //评论来源
	Commenter string `xml:"commenter" json:"commenter"` //评论者
}
