package createpnr

import (
	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	unirs "github.com/redochen/demos/travelport-uapi/soap/universal/response"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//CreatePnrResult 创建PNR结果
type CreatePnrResult struct {
	models.BaseResult
	UniversalRecord   *UniversalRecord          `xml:"universalRecord" json:"universalRecord"`     //全局记录
	Travelers         []*Traveler               `xml:"travelers" json:"travelers"`                 //旅客列表
	OtherServiceInfos []*OtherServiceInfo       `xml:"otherServiceInfo" json:"otherServiceInfo"`   //其他服务信息列表
	ProviderInfos     []*ProviderInfo           `xml:"providerInfos" json:"providerInfos"`         //供应商预订列表
	AirReservations   []*AirReservationResult   `xml:"airReservations" json:"airReservations"`     //机票预订列表
	HotelReservations []*HotelReservationResult `xml:"hotelReservations" json:"hotelReservations"` //酒店预订列表
	AgencyInfos       []*AgencyInfo             `xml:"agencyInfos" json:"agencyInfos"`             //代理商信息
	Payments          []*FormOfPayment          `xml:"payments" json:"payments"`                   //支付类型列表
}

//UniversalRecord 全局记录类
type UniversalRecord struct {
	Code    string `xml:"code" json:"code"`       //UR编号
	Version string `xml:"version" json:"version"` //UR版本号
	Status  string `xml:"status" json:"status"`   //状态
}

//AgencyInfo 代理商信息
type AgencyInfo struct {
	ActionType string `xml:"actionType" json:"actionType"` //事件类型
	AgentCode  string `xml:"agentCode" json:"agentCode"`   //机构代码
	BranchCode string `xml:"branchCode" json:"branchCode"` //分支代码
	AgencyCode string `xml:"agencyCode" json:"agencyCode"` //代理人代码
	EventTime  string `xml:"eventTime" json:"eventTime"`   //事件发生的时间
}

//ProviderInfo 供应商预订信息
type ProviderInfo struct {
	ProviderCode   string `xml:"providerCode" json:"providerCode"`         //供应商代码：1G
	LocatorCode    string `xml:"locatorCode" json:"locatorCode"`           //PNR编码
	CreateDate     string `xml:"createDate" json:"createDate"`             //创建时间
	ModifiedDate   string `xml:"modifiedDate" json:"modifiedDate"`         //修改时间
	HostCreateDate string `xml:"hostCreateDate" json:"hostCreateDate"`     //创建时间
	OwningPCC      string `xml:"owningPCC" json:"owningPCC"`               //所属PCC
	Key            string `xml:"key,omitempty" json:"key,omitempty"`       //引用键
	ElStat         string `xml:"ElStat,omitempty" json:"ElStat,omitempty"` //状态
}

//ReservationResult 预订结果类
type ReservationResult struct {
	LocatorCode        string   `xml:"locatorCode" json:"locatorCode"`               //预订编号
	CreateDate         string   `xml:"createDate" json:"createDate"`                 //创建时间
	ModifiedDate       string   `xml:"modifiedDate" json:"modifiedDate"`             //修改时间
	BookingTravelerKey []string `xml:"bookingTravelerKey" json:"bookingTravelerKey"` //预订旅客引用KEY
}

//AirReservationResult 机票预订结果类
type AirReservationResult struct {
	ReservationResult
	Suppliers []*SupplierResult `xml:"suppliers" json:"suppliers"` //供应商列表
	Segments  []*Segment        `xml:"segments" json:"segments"`   //航段列表
}

//SupplierResult 供应商结果类
type SupplierResult struct {
	Code               string `xml:"code" json:"code"`                             //供应商代码：MU
	LocatorCode        string `xml:"locatorCode" json:"locatorCode"`               //预订编号
	ReservationInfoRef string `xml:"reservationInfoRef" json:"reservationInfoRef"` //预订信息索引
	CreateDate         string `xml:"createDate" json:"createDate"`                 //创建时间
}

//HotelReservationResult 酒店预订结果
type HotelReservationResult struct {
	ReservationResult
	TravelerKey  string         `xml:"travelerKey" json:"travelerKey"`               //旅客引用键
	Surname      string         `xml:"surname" json:"surname"`                       //旅客姓氏
	GivenName    string         `xml:"givenName" json:"givenName"`                   //旅客名字
	Property     *HotelProperty `xml:"property" json:"property"`                     //酒店物业
	Rates        []*HotelRate   `xml:"rates" json:"rates"`                           //价格列表
	Rooms        int            `xml:"rooms,omitempty" json:"rooms,omitempty"`       //房间数量，默认为1
	Adults       int            `xml:"adults,omitempty" json:"adults,omitempty"`     //成人数量，默认为1
	Children     int            `xml:"children,omitempty" json:"children,omitempty"` //儿童数量，默认为0
	CheckinDate  string         `xml:"checkin" json:"checkin"`                       //入住日期，格式为“yyyy-MM-dd”
	CheckoutDate string         `xml:"checkout" json:"checkout"`                     //退房日期，格式为“yyyy-MM-dd”
	SourceCode   string         `xml:"sourceCode" json:"sourceCode"`                 //预订来源代码
	SourceType   string         `xml:"sourceType" json:"sourceType"`                 //预订来源类型
	SellMessage  []string       `xml:"sellMessage" json:"sellMessage"`               //售卖信息
}

//SetErrorCode 设置错误代码
func SetErrorCode(code ErrorCode) *CreatePnrResult {
	result := &CreatePnrResult{}
	result.SetErrorCode(code)
	return result
}

//SetErrorCode 设置错误代码
func (createPnrResult *CreatePnrResult) SetErrorCode(code ErrorCode) *CreatePnrResult {
	createPnrResult.BaseResult.SetErrorCode(code)
	return createPnrResult
}

//SetErrorMessage 设置错误消息
func SetErrorMessage(message string) *CreatePnrResult {
	result := &CreatePnrResult{}
	result.SetErrorMessage(message)
	return result
}

//SetErrorMessage 设置错误消息
func (createPnrResult *CreatePnrResult) SetErrorMessage(message string) *CreatePnrResult {
	createPnrResult.BaseResult.SetErrorMessage(message)
	return createPnrResult
}

//ParsePnrCode 解析PNR编码
func (createPnrResult *CreatePnrResult) ParsePnrCode(record *unirs.UniversalRecord) {
	if nil == createPnrResult || nil == record {
		return
	}

	//解析大编码
	createPnrResult.UniversalRecord = &UniversalRecord{
		Code:    string(record.LocatorCode),
		Version: string(record.Version),
		Status:  record.Status,
	}

	//解析小编码（订座终端里的编码）
	if record.ProviderReservationInfo != nil && len(record.ProviderReservationInfo) > 0 {
		for _, info := range record.ProviderReservationInfo {
			provider := &ProviderInfo{
				Key:            string(info.Key),
				ProviderCode:   string(info.ProviderCode),
				LocatorCode:    string(info.LocatorCode),
				CreateDate:     info.CreateDate,
				ModifiedDate:   info.ModifiedDate,
				HostCreateDate: info.HostCreateDate,
				OwningPCC:      string(info.OwningPCC),
			}
			createPnrResult.ProviderInfos = append(createPnrResult.ProviderInfos, provider)
		}
	}
}
