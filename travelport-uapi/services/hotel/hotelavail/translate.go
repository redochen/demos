package hotelavail

import (
	"errors"
	. "github.com/redochen/demos/travelport-uapi/models/hotel"
	. "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavail"
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	hot "github.com/redochen/demos/travelport-uapi/soap/hotel"
	hotproxy "github.com/redochen/demos/travelport-uapi/soap/hotel/proxy"
	hotrq "github.com/redochen/demos/travelport-uapi/soap/hotel/request"
	hotrs "github.com/redochen/demos/travelport-uapi/soap/hotel/response"
	. "github.com/redochen/tools/time"
)

//getReqEnvolpe 获取请求参数
func getReqEnvolpe(param *HotelAvailParam, branchCode string) (*soap.ReqEnvelope, error) {
	if nil == param {
		return nil, errors.New("param cannot be nil")
	}

	if param.CheckinDate == "" || param.CheckoutDate == "" {
		return nil, errors.New("Check-in and Check-out dates are required")
	}

	body := hotproxy.NewHotelAvailReqBody(branchCode)
	if nil == body {
		return nil, errors.New("failed to create HotelAvailReqBody")
	}

	req := body.Request

	//设置地点
	req.HotelSearchLocation = &hotrq.HotelSearchLocation{
		HotelLocation: &hotrq.HotelLocation{
			Location: com.TypeIATACode(param.Location),
		},
	}

	if param.LocationType == 1 {
		req.HotelSearchLocation.HotelLocation.LocationType = "City"
	} else if param.LocationType == 2 {
		req.HotelSearchLocation.HotelLocation.LocationType = "Airport"
	}

	if len(param.Reference) > 0 {
		req.HotelSearchLocation.ReferencePoint = &hotrq.HotelReferencePoint{
			Value: com.TypeReferencePoint(param.Reference),
		}
	}

	//设置时间
	req.HotelStay = &hotrq.HotelStay{
		CheckinDate:  hot.TypeDate(CcTime.AddDateSeparator(param.CheckinDate, "-", true)),
		CheckoutDate: hot.TypeDate(CcTime.AddDateSeparator(param.CheckoutDate, "-", true)),
	}

	//获取查询选项
	req.HotelSearchModifiers = getHotelSearchModifiers(param)

	return soap.NewReqEnvelope(body), nil
}

//getHotelSearchModifiers 获取HotelSearchModifiers
func getHotelSearchModifiers(param *HotelAvailParam) *hotrq.HotelSearchModifiers {
	if nil == param {
		return nil
	}

	//设置人数
	modifiers := &hotrq.HotelSearchModifiers{
		NumberOfAdults: param.Adults,
	}

	//指定1G
	modifiers.PermittedProviders = &comrq.PermittedProviders{
		Provider: &comrq.Provider{
			Code: param.ProviderCode, //"1G",
		},
	}

	if param.Children > 0 {
		modifiers.NumberOfChildren = &hotrq.NumberOfChildren{
			Count: param.Children,
		}
	}

	//设置房间数
	if param.Rooms > 0 {
		modifiers.NumberOfRooms = param.Rooms
	}

	//设置酒店名称
	if len(param.Name) > 0 {
		modifiers.HotelName = param.Name
	}

	//设置星级
	if len(param.RatingType) > 0 {
		modifiers.HotelRating = make([]*hotrq.HotelRating, 0)

		rating := &hotrq.HotelRating{
			RatingProvider: param.RatingType,
		}

		if param.Ratings != nil && len(param.Ratings) > 0 {
			rating.Rating = make([]hot.TypeSimpleHotelRating, 0)
			for _, r := range param.Ratings {
				rating.Rating = append(rating.Rating, hot.TypeSimpleHotelRating(r))
			}

		} else {
			rating.RatingRange = &hotrq.RatingRange{
				MinimumRating: hot.TypeSimpleHotelRating(param.MinRating),
				MaximumRating: hot.TypeSimpleHotelRating(param.MaxRating),
			}
		}

		modifiers.HotelRating = append(
			modifiers.HotelRating,
			rating)
	}

	//设置包含的酒店
	if param.IncludeHotels != nil && len(param.IncludeHotels) > 0 {
		modifiers.PermittedChains = &hotrq.HotelChains{
			HotelChain: make([]*hotrq.HotelChain, 0),
		}

		for _, ih := range param.IncludeHotels {
			ic := &hotrq.HotelChain{
				Code: com.TypeHotelChainCode(ih),
			}

			modifiers.PermittedChains.HotelChain = append(
				modifiers.PermittedChains.HotelChain,
				ic)
		}
	}

	//设置排除的酒店
	if param.ExcludeHotels != nil && len(param.ExcludeHotels) > 0 {
		modifiers.ProhibitedChains = &hotrq.HotelChains{
			HotelChain: make([]*hotrq.HotelChain, 0),
		}
		for _, eh := range param.ExcludeHotels {
			ec := &hotrq.HotelChain{
				Code: com.TypeHotelChainCode(eh),
			}

			modifiers.ProhibitedChains.HotelChain = append(
				modifiers.ProhibitedChains.HotelChain,
				ec)
		}
	}

	//设置价格类别
	if param.Categories != nil && len(param.Categories) > 0 {
		modifiers.RateCategory = make([]com.TypeOTACode, 0)
		for _, rc := range param.Categories {
			if rc <= 0 {
				continue
			}

			modifiers.RateCategory = append(
				modifiers.RateCategory,
				com.TypeOTACode(rc))
		}
	}

	//设置房间特性
	if param.Amenities != nil && len(param.Amenities) > 0 {
		modifiers.Amenities = &hotrq.Amenities{
			Amenity: make([]*hotrq.Amenity, 0),
		}

		for _, a := range param.Amenities {
			if a <= 0 {
				continue
			}

			modifiers.Amenities.Amenity = append(
				modifiers.Amenities.Amenity,
				&hotrq.Amenity{
					Code: hot.TypeAmenity(a),
				})

		}
	}

	//设置货币代码
	if len(param.Currency) > 0 {
		modifiers.PreferredCurrency = com.TypeCurrency(param.Currency)
	}

	return modifiers
}

//getResult 转换结果
func getResult(body *hotproxy.HotelAvailRspBody) (*HotelAvailResult, error) {
	if nil == body || nil == body.Response {
		return nil, errors.New("response is empty")
	}

	rsp := body.Response
	if nil == rsp {
		return nil, errors.New("HotelSearchAvailabilityRsp is empty")
	}

	result := &HotelAvailResult{}

	//解析HostToken
	if rsp.HostToken != nil {
		result.HostToken = &HostTokenResult{
			Host:  string(rsp.HostToken.Host),
			Key:   rsp.HostToken.Key,
			Token: rsp.HostToken.Value,
		}
	}

	//解析酒店信息
	if rsp.HotelSearchResult != nil && len(rsp.HotelSearchResult) > 0 {
		result.Hotels = make([]*HotelInfoResult, 0)
		for _, hsr := range rsp.HotelSearchResult {
			info := translateHotelSearchResultToHotelInfo(hsr)
			if info != nil {
				result.Hotels = append(result.Hotels, info)
			}
		}
	}

	return result, nil
}

//translateHotelSearchResultToHotelInfo 将HotelSearchResult转换为HotelInfo
func translateHotelSearchResultToHotelInfo(result *hotrs.HotelSearchResult) *HotelInfoResult {
	if nil == result {
		return nil
	}

	info := &HotelInfoResult{}

	//解析供应商信息
	if result.VendorLocation != nil && len(result.VendorLocation) > 0 {
		info.Vendors = make([]*VendorInfoResult, 0)
		for _, vl := range result.VendorLocation {
			vendor := translateVendorLocationToVendorInfo(vl)
			if vendor != nil {
				info.Vendors = append(info.Vendors, vendor)
			}
		}
	}

	//解析物业信息
	if result.HotelProperty != nil && len(result.HotelProperty) > 0 {
		info.Properties = make([]*HotelProperty, 0)
		for _, hp := range result.HotelProperty {
			property := TranslateHotelPropertyResult(hp)
			if property != nil {
				info.Properties = append(info.Properties, property)
			}
		}
	}

	//解析折扣信息
	if result.CorporateDiscountID != nil && len(result.CorporateDiscountID) > 0 {
		info.Discounts = make([]*CorporateDiscount, 0)
		for _, cd := range result.CorporateDiscountID {
			discount := translateCorporateDiscount(cd)
			if discount != nil {
				info.Discounts = append(info.Discounts, discount)
			}
		}
	}

	//解析价格信息
	if result.RateInfo != nil && len(result.RateInfo) > 0 {
		info.Rates = make([]*RateInfoResult, 0)
		for _, ri := range result.RateInfo {
			rate := translateRateInfo(ri)
			if rate != nil {
				info.Rates = append(info.Rates, rate)
			}
		}
	}

	//解析媒介信息
	if result.MediaItem != nil {
		info.Media = &MediaInfoResult{
			Caption:  result.MediaItem.Caption,
			Height:   result.MediaItem.Height,
			Width:    result.MediaItem.Width,
			Type:     result.MediaItem.Type,
			Url:      result.MediaItem.Url,
			Icon:     result.MediaItem.Icon,
			SizeCode: result.MediaItem.SizeCode,
		}
	}

	if result.HotelType != nil {
		info.SourceLink = bool(result.HotelType.SourceLink)
	}

	if result.PropertyDescription != nil {
		info.Description = result.PropertyDescription.Value
		info.ProviderCode = string(result.PropertyDescription.ProviderCode)
	}

	return info
}

//translateVendorLocationToVendorInfo 将VendorLocation转换为VendorInfo
func translateVendorLocationToVendorInfo(vendor *comrs.VendorLocation) *VendorInfoResult {
	if nil == vendor {
		return nil
	}

	info := &VendorInfoResult{
		Key:              string(vendor.Key),
		ProviderCode:     string(vendor.ProviderCode),
		VendorCode:       string(vendor.VendorCode),
		VendorLocationID: vendor.VendorLocationID,
	}

	return info
}

//TranslateHotelPropertyResult 转换HotelProperty信息
func TranslateHotelPropertyResult(hotel *hotrs.HotelProperty) *HotelProperty {
	if nil == hotel {
		return nil
	}

	property := &HotelProperty{
		Code:                  string(hotel.HotelCode),
		Name:                  hotel.Name,
		Chain:                 string(hotel.HotelChain),
		Location:              string(hotel.HotelLocation),
		Transportation:        string(hotel.HotelTransportation),
		ReserveRequirement:    string(hotel.ReserveRequirement),
		Availability:          hotel.Availability,
		PreferredOption:       hotel.PreferredOption,
		ParticipationLevel:    string(hotel.ParticipationLevel),
		NetTransCommissionInd: hotel.NetTransCommissionInd,
		Key:                   string(hotel.Key),
		VendorLocationKey:     hotel.VendorLocationKey,
	}

	//解析地址信息
	if hotel.PropertyAddress != nil && hotel.PropertyAddress.Address != nil &&
		len(hotel.PropertyAddress.Address) > 0 {
		property.Address = make([]string, 0)
		for _, addr := range hotel.PropertyAddress.Address {
			if addr != "" {
				property.Address = append(property.Address, addr)
			}
		}
	}

	//解析电话信息
	if hotel.PhoneNumber != nil && len(hotel.PhoneNumber) > 0 {
		property.Phones = make([]*PhoneInfo, 0)
		for _, pn := range hotel.PhoneNumber {
			phone := translatePhoneNumberToPhoneInfo(pn)
			if phone != nil {
				property.Phones = append(property.Phones, phone)
			}
		}
	}

	//解析酒店坐标
	if hotel.CoordinateLocation != nil {
		property.Latitude = hotel.CoordinateLocation.Latitude
		property.Longitude = hotel.CoordinateLocation.Longitude
	}

	//解析距离信息
	if hotel.Distance != nil {
		property.DistanceUnits = hotel.Distance.Units
		property.DistanceValue = hotel.Distance.Value
		property.DistanceDirection = hotel.Distance.Direction
	}

	//解析评级信息
	if hotel.HotelRating != nil && len(hotel.HotelRating) > 0 {
		property.Ratings = make([]*RatingInfo, 0)
		for _, hr := range hotel.HotelRating {
			rating := translateHotelRatingToRatingInfo(hr)
			if rating != nil {
				property.Ratings = append(property.Ratings, rating)
			}
		}
	}

	//解析酒店设施
	if hotel.Amenities != nil && hotel.Amenities.Amenity != nil &&
		len(hotel.Amenities.Amenity) > 0 {
		property.Amenities = make([]string, 0)
		for _, a := range hotel.Amenities.Amenity {
			property.Amenities = append(property.Amenities, string(a.Code))
		}
	}

	return property
}

//translatePhoneNumberToPhoneInfo 将PhoneNumber转换为PhoneInfo
func translatePhoneNumberToPhoneInfo(phone *comrs.PhoneNumber) *PhoneInfo {
	if nil == phone {
		return nil
	}

	info := &PhoneInfo{
		Key:         string(phone.Key),
		Type:        phone.Type,
		Location:    phone.Location,
		CountryCode: phone.CountryCode,
		AreaCode:    phone.AreaCode,
		Number:      phone.Number,
		Extension:   phone.Extension,
		Text:        phone.Text,
	}

	return info
}

//translateHotelRatingToRatingInfo 将HotelRating转换为RatingInfo
func translateHotelRatingToRatingInfo(rating *hotrs.HotelRating) *RatingInfo {
	if nil == rating {
		return nil
	}

	info := &RatingInfo{
		Type:    rating.RatingProvider,
		Ratings: make([]int, 0),
	}

	if rating.Rating != nil && len(rating.Rating) > 0 {
		for _, r := range rating.Rating {
			info.Ratings = append(info.Ratings, int(r))
		}

	} else if rating.RatingRange != nil {
		for i := rating.RatingRange.MinimumRating; i < rating.RatingRange.MaximumRating; i++ {
			info.Ratings = append(info.Ratings, int(i))
		}
	}

	return info
}

//translateCorporateDiscount 转换CorporateDiscount
func translateCorporateDiscount(discount *comrs.CorporateDiscountID) *CorporateDiscount {
	if nil == discount {
		return nil
	}

	result := &CorporateDiscount{
		Value:                discount.Value,
		IsNegotiatedRateCode: discount.NegotiatedRateCode,
	}

	return result
}

//translateRateInfo 转换RateInfo
func translateRateInfo(rate *hotrs.RateInfo) *RateInfoResult {
	if nil == rate {
		return nil
	}

	info := &RateInfoResult{
		MinAmount:                        rate.MinimumAmount.GetFloatAmount(),
		MinAmountCurrency:                rate.MinimumAmount.GetCurrency(),
		ApproximateMinAmount:             rate.ApproximateMinimumAmount.GetFloatAmount(),
		ApproximateMinAmountCurrency:     rate.ApproximateMinimumAmount.GetCurrency(),
		MinAmountRateChanged:             rate.MinAmountRateChanged,
		MaxAmount:                        rate.MaximumAmount.GetFloatAmount(),
		MaxAmountCurrency:                rate.MaximumAmount.GetCurrency(),
		ApproximateMaxAmount:             rate.ApproximateMaximumAmount.GetFloatAmount(),
		ApproximateMaxAmountCurrency:     rate.ApproximateMaximumAmount.GetCurrency(),
		MaxAmountRateChanged:             rate.MaxAmountRateChanged,
		MinStayAmount:                    rate.MinimumStayAmount.GetFloatAmount(),
		MinStayAmountCurrency:            rate.MinimumStayAmount.GetCurrency(),
		ApproximateMinStayAmount:         rate.ApproximateMinimumStayAmount.GetFloatAmount(),
		ApproximateMinStayAmountCurrency: rate.ApproximateMinimumStayAmount.GetCurrency(),
		Commission:                       rate.Commission,
		RateSupplier:                     string(rate.RateSupplier),
		RateSupplierLogo:                 rate.RateSupplierLogo,
		PaymentType:                      rate.PaymentType,
	}

	return info
}
