package air

//Price 价格元素类
type Price struct {
	BasePrice         float32 `xml:"basePrice" json:"basePrice"`                               //基础价
	BaseCurrency      string  `xml:"baseCurrency" json:"baseCurrency"`                         //基础价货币代码
	AppBasePrice      float32 `xml:"appBasePrice,omitempty" json:"appBasePrice,omitempty"`     //近似基础价
	AppBaseCurrency   string  `xml:"appBaseCurrency" json:"appBaseCurrency"`                   //近似基础价货币代码
	EquivBasePrice    float32 `xml:"equivBasePrice,omitempty" json:"equivBasePrice,omitempty"` //等效基础价
	EquivBaseCurrency string  `xml:"equivBaseCurrency" json:"equivBaseCurrency"`               //等效基础价货币代码
	Taxes             float32 `xml:"taxes,omitempty" json:"taxes,omitempty"`                   //税费
	TaxesCurrency     string  `xml:"taxesCurrency" json:"taxesCurrency"`                       //税费货币代码
	AppTaxes          float32 `xml:"appTaxes,omitempty" json:"appTaxes,omitempty"`             //近似税费
	AppTaxesCurrency  string  `xml:"appTaxesCurrency" json:"appTaxesCurrency"`                 //近似税费货币代码
	TotalPrice        float32 `xml:"totalPrice,omitempty" json:"totalPrice,omitempty"`         //总价
	TotalCurrency     string  `xml:"totalCurrency" json:"totalCurrency"`                       //总价货币代码
	AppTotalPrice     float32 `xml:"appTotalPrice,omitempty" json:"appTotalPrice,omitempty"`   //近似总价
	AppTotalCurrency  string  `xml:"appTotalCurrency" json:"appTotalCurrency"`                 //近似总价货币代码
}

//PricingInfo 价格信息类
type PricingInfo struct {
	PricePointKey         string         `xml:"pricePointKey" json:"pricePointKey"`                       //AirPricePoint键
	PricingInfoKey        string         `xml:"pricingInfoKey" json:"pricingInfoKey"`                     //AirPricingInfo键
	PassengerType         string         `xml:"passengerType" json:"passengerType"`                       //乘客类型
	Price                 *Price         `xml:"price,omitempty" json:"price,omitempty"`                   //价格信息
	PlatingCarrier        string         `xml:"platingCarrier" json:"platingCarrier"`                     //出票航司
	LatestTicketingTime   string         `xml:"latestTicketingTime" json:"latestTicketingTime"`           //最后出票时间
	PrivateFareType       string         `xml:"privateFare" json:"privateFare"`                           //私有运价类型
	FareCalc              string         `xml:"fareCalc" json:"fareCalc"`                                 //NUC值
	FareInfos             []*FareInfo    `xml:"fareInfos,omitempty" json:"fareInfos,omitempty"`           //运价信息列表
	TaxInfos              []*TaxInfo     `xml:"taxInfos,omitempty" json:"taxInfos,omitempty"`             //税费列表
	PricingMethod         string         `xml:"pricingMethod,omitempty" json:"pricingMethod,omitempty"`   //计价方式
	Exchangeable          bool           `xml:"exchangeable,omitempty" json:"exchangeable,omitempty"`     //是否可改期
	ChangePenaltyAmount   string         `xml:"changeAmount,omitempty" json:"changeAmount,omitempty"`     //改期费用金额
	ChangePenaltyCurrency string         `xml:"changeCurrency,omitempty" json:"changeCurrency,omitempty"` //改期费用货币代码
	ChangePenaltyPercent  string         `xml:"changePercent,omitempty" json:"changePercent,omitempty"`   //改期费用百分比
	Refundable            bool           `xml:"refundable,omitempty" json:"refundable,omitempty"`         //是否可退票
	CancelPenaltyAmount   string         `xml:"cancelAmount,omitempty" json:"cancelAmount,omitempty"`     //退票费用金额
	CancelPenaltyCurrency string         `xml:"cancelCurrency,omitempty" json:"cancelCurrency,omitempty"` //退票费用货币代码
	CancelPenaltyPercent  string         `xml:"cancelPercent,omitempty" json:"cancelPercent,omitempty"`   //退票费用百分比
	ProviderCode          string         `xml:"providerCode,omitempty" json:"providerCode,omitempty"`     //供应商代码
	IncludesVAT           bool           `xml:"includesVAT,omitempty" json:"includesVAT,omitempty"`       //
	BaggageInfos          []*BaggageInfo `xml:"baggageInfos,omitempty" json:"baggageInfos,omitempty"`     //详细的行李额规定
}

//FareInfo 运价信息类
type FareInfo struct {
	FareInfoKey     string  `xml:"fareInfoKey" json:"fareInfoKey"`         //FareInfo键
	FareBasis       string  `xml:"fareBasis" json:"fareBasis"`             //运价基础
	PassengerType   string  `xml:"passengerType" json:"passengerType"`     //乘客类型
	Amount          float32 `xml:"amount" json:"amount"`                   //金额
	Currency        string  `xml:"currency" json:"currency"`               //货币代码
	Origin          string  `xml:"origin" json:"origin"`                   //起始地
	Destination     string  `xml:"destination" json:"destination"`         //目的地
	DepartureDate   string  `xml:"departure" json:"departure"`             //出发日期
	NotValidBefore  string  `xml:"notValidBefore" json:"notValidBefore"`   //有效开始日期
	NotValidAfter   string  `xml:"notValidAfter" json:"notValidAfter"`     //有效结束日期
	EffectiveDate   string  `xml:"effectiveDate" json:"effectiveDate"`     //生效日期
	PrivateFareType string  `xml:"privateFareType" json:"privateFareType"` //私有运价类型
}

//TaxInfo 税费信息类
type TaxInfo struct {
	TaxInfoKey string  `xml:"taxInfoKey" json:"taxInfoKey"` //TaxInfo键
	Category   string  `xml:"category" json:"category"`     //类别
	Amount     float32 `xml:"amount" json:"amount"`         //金额
	Currency   string  `xml:"currency" json:"currency"`     //货币代码
}

//BaggageInfo 行李信息类
type BaggageInfo struct {
	PassengerType  string            `xml:"passengerType" json:"passengerType"`                       //旅客类型
	Origin         string            `xml:"origin" json:"origin"`                                     //出发地
	Destination    string            `xml:"destination" json:"destination"`                           //目的地
	NumberOfPieces int               `xml:"numberOfPieces,omitempty" json:"numberOfPieces,omitempty"` //件数
	MaxWeight      *BaggageMeasure   `xml:"maxWeight,omitempty" json:"maxWeight,omitempty"`           //最大重量
	IsCarryOn      bool              `xml:"isCarryOn,omitempty" json:"isCarryOn,omitempty"`           //是否为托运信息
	Carrier        string            `xml:"carrier,omitempty" json:"carrier,omitempty"`               //航司
	Details        []*BaggageDetails `xml:"details,omitempty" json:"details,omitempty"`               //详情
	Text           []string          `xml:"text" json:"text"`                                         //文本信息
}

//BaggageDetails 行李详情
type BaggageDetails struct {
	Restriction    []*BaggageRestriction `xml:"restriction,omitempty" json:"restriction,omitempty"` //限制信息
	ApplicableInfo string                `xml:"applicableInfo" json:"applicableInfo"`               //适用信息
	Price          *Price                `xml:"price,omitempty" json:"price,omitempty"`             //价格信息
}

//BaggageRestriction 行李限制
type BaggageRestriction struct {
	MaxWeight  *BaggageMeasure   `xml:"maxWeight,omitempty" json:"maxWeight,omitempty"`   //最大重量
	Dimensions []*BaggageMeasure `xml:"dimensions,omitempty" json:"dimensions,omitempty"` //行李体积
	Text       []string          `xml:"text" json:"text"`                                 //文本信息
}

//BaggageMeasure 行李测量
type BaggageMeasure struct {
	Type  string  `xml:"type" json:"type"`                       //类型：长、宽、高、重量
	Value float32 `xml:"value,omitempty" json:"value,omitempty"` //值
	Unit  string  `xml:"unit" json:"unit"`                       //单位
}
