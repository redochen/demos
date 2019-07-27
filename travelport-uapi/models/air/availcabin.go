package air

//AvailInfo 可用舱位集
type AvailInfo struct {
	ProviderCode string        `xml:"providerCode" json:"providerCode"` //供应商代码
	AvailCabins  []*AvailCabin `xml:"availCabins" json:"availCabins"`   //可用舱位列表
}

//AvailCabin 可用舱位
type AvailCabin struct {
	CabinClass  string `xml:"cabinClass" json:"cabinClass"`   //舱位等级
	BookingCode string `xml:"bookingCode" json:"bookingCode"` //舱位
	AvailCount  string `xml:"availCount" json:"availCount"`   //可用数量
}
