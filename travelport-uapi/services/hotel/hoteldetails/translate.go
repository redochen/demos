package hoteldetails

import (
	"errors"
	cfg "github.com/redochen/demos/travelport-uapi/config"
	. "github.com/redochen/demos/travelport-uapi/models/hotel"
	hotavmodel "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavail"
	. "github.com/redochen/demos/travelport-uapi/models/hotel/hoteldetails"
	hotavsvc "github.com/redochen/demos/travelport-uapi/services/hotel/hotelavail"
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
func getReqEnvolpe(param *HotelDetailsParam, branchCode string) (*soap.ReqEnvelope, error) {
	if nil == param {
		return nil, errors.New("param cannot be nil")
	}

	if param.Code == "" {
		return nil, errors.New("hotel code is required")
	}

	body := hotproxy.NewHotelDetailsReqBody(branchCode)
	if nil == body {
		return nil, errors.New("failed to create HotelMediaReqBody")
	}

	req := body.Request

	//设置酒店属性
	req.HotelProperty = &hotrq.HotelProperty{
		HotelChain: com.TypeHotelChainCode(param.Chain),
		HotelCode:  com.TypeHotelCode(param.Code),
		Name:       param.Name,
	}

	//获取查询选项
	req.HotelDetailsModifiers = getHotelDetailsModifiers(param)

	//设置其他选项
	req.ReturnMediaLinks = true
	req.ReturnGuestReviews = true

	//设置主机令牌
	if param.HostToken != nil {
		req.HostToken = &comrq.HostToken{
			Value: param.HostToken.Token,
			Host:  com.TypeProviderCode(param.HostToken.Host),
			Key:   param.HostToken.Key,
		}
	}

	return soap.NewReqEnvelope(body), nil
}

//getHotelDetailsModifiers 获取置酒店选项
func getHotelDetailsModifiers(param *HotelDetailsParam) *hotrq.HotelDetailsModifiers {
	modifiers := &hotrq.HotelDetailsModifiers{
		MaxWait: uint(cfg.TimeoutSeconds * 1000),
	}

	//指定1G
	modifiers.PermittedProviders = &comrq.PermittedProviders{
		Provider: &comrq.Provider{
			Code: param.ProviderCode, //"1G",
		},
	}

	if nil == param {
		return modifiers
	}

	//设置人数
	if param.Adults > 0 {
		modifiers.NumberOfAdults = param.Adults
	}

	//设置房间数
	if param.Rooms > 0 {
		modifiers.NumberOfRooms = param.Rooms
	}

	//设置日期
	if param.CheckinDate != "" {
		modifiers.HotelStay = &hotrq.HotelStay{
			CheckinDate:  hot.TypeDate(CcTime.AddDateSeparator(param.CheckinDate, "-", true)),
			CheckoutDate: hot.TypeDate(CcTime.AddDateSeparator(param.CheckoutDate, "-", true)),
		}
	}

	//设置货币代码
	if param.Currency != "" {
		modifiers.PreferredCurrency = com.TypeCurrency(param.Currency)
	}

	//设置接口模式
	if param.OnlyDescription != 0 {
		modifiers.RateRuleDetail = "None"
	} else {
		modifiers.RateRuleDetail = "Complete"
	}

	//设置供应商
	if param.Provider != "" {
		modifiers.PermittedProviders = &comrq.PermittedProviders{
			Provider: &comrq.Provider{
				Code: param.Provider,
			},
		}
	} /*else {
		modifiers.PermittedProviders = &comrq.PermittedProviders{
			Provider: &comrq.Provider{
				Code: "TRM",
			},
		}
	}*/

	//设置第三方供应商
	if param.Aggregator != "" {
		modifiers.PermittedAggregators = &hotrq.PermittedAggregators{
			Aggregator: make([]*hotrq.Aggregator, 0),
		}

		//Note that Supplier in this usage indicates the TRM aggregator,
		//not the hotel chain or property supplier.
		modifiers.PermittedAggregators.Aggregator = append(
			modifiers.PermittedAggregators.Aggregator,
			&hotrq.Aggregator{
				Name: com.TypeThirdPartySupplier(param.Aggregator),
			})
	}

	return modifiers
}

//getResult 转换结果
func getResult(body *hotproxy.HotelDetailsRspBody) (*HotelDetailsResult, error) {
	if nil == body || nil == body.Response {
		return nil, errors.New("response is empty")
	}

	rsp := body.Response
	if nil == rsp {
		return nil, errors.New("HotelDetailsRsp is empty")
	}

	result := &HotelDetailsResult{}

	if rsp.RequestedHotelDetails != nil {
		parseRequestedHotelDetails(result, rsp.RequestedHotelDetails)
	} else if rsp.HotelAlternateProperties != nil {
		parseHotelAlternateProperties(result, rsp.HotelAlternateProperties)
	}

	//解析GuestReviews
	if rsp.GuestReviews != nil && rsp.GuestReviews.Comments != nil &&
		len(rsp.GuestReviews.Comments) > 0 {
		result.Comments = make([]*Comments, 0)

		for _, grc := range rsp.GuestReviews.Comments {
			c := &Comments{
				Content:   grc.Value,
				Date:      grc.Date,
				Language:  string(grc.CommenterLanguage),
				Source:    grc.CommentSourceName,
				Commenter: grc.Commenter,
			}

			result.Comments = append(result.Comments, c)
		}
	}

	return result, nil
}

//parseRequestedHotelDetails 解析RequestedHotelDetails
func parseRequestedHotelDetails(result *HotelDetailsResult, details *hotrs.RequestedHotelDetails) {
	if nil == result || nil == details {
		return
	}

	//解析酒店物业
	if details.HotelProperty != nil {
		result.Property = hotavsvc.TranslateHotelPropertyResult(details.HotelProperty)
	}

	//解析酒店详情
	if details.HotelDetailItem != nil && len(details.HotelDetailItem) > 0 {
		result.Items = make([]*HotelDetailItem, 0)
		for _, hdi := range details.HotelDetailItem {
			item := &HotelDetailItem{
				Name: hdi.Name,
				Text: hdi.Text,
			}

			result.Items = append(result.Items, item)
		}
	}

	//解析价格详情
	if details.HotelRateDetail != nil && len(details.HotelRateDetail) > 0 {
		result.Rates = make([]*HotelRateDetail, 0)
		for _, hrd := range details.HotelRateDetail {
			rate := TranslateHotelRateDetailResult(hrd)
			if rate != nil {
				result.Rates = append(result.Rates, rate)
			}
		}
	}

	//解析MediaItem
	if details.MediaItem != nil && len(details.MediaItem) > 0 {
		result.Medias = make([]*hotavmodel.MediaInfoResult, 0)
		for _, mi := range details.MediaItem {
			media := translateMediaItemResult(mi)
			if media != nil {
				result.Medias = append(result.Medias, media)
			}
		}
	}

	//解析HotelType
	if details.HotelType != nil {
		result.SourceLink = bool(details.HotelType.SourceLink)
	}
}

//parseHotelAlternateProperties 解析HotelAlternateProperties
func parseHotelAlternateProperties(result *HotelDetailsResult, properties *hotrs.HotelAlternateProperties) {
	if nil == result || nil == properties {
		return
	}

	if nil == properties.HotelProperty || len(properties.HotelProperty) == 0 {
		return
	}

	//解析酒店物业
	result.AlternateProperties = make([]*HotelProperty, 0)

	for _, p := range properties.HotelProperty {
		property := hotavsvc.TranslateHotelPropertyResult(p)
		if property != nil {
			result.AlternateProperties = append(result.AlternateProperties, property)
		}
	}
}

//TranslateHotelRateDetailResult 转换HotelRateDetail
func TranslateHotelRateDetailResult(detail *hotrs.HotelRateDetail) *HotelRateDetail {
	if nil == detail {
		return nil
	}

	result := &HotelRateDetail{
		SupplementalRateInfo: detail.SupplementalRateInfo,
		RatePlanType:         string(detail.RatePlanType),
		Rate: &RateInfo{
			Base:              detail.Base.GetFloatAmount(),
			BaseCurrency:      detail.Base.GetCurrency(),
			Tax:               detail.Tax.GetFloatAmount(),
			TaxCurrency:       detail.Tax.GetCurrency(),
			Total:             detail.Total.GetFloatAmount(),
			TotalCurrency:     detail.Total.GetCurrency(),
			Surcharge:         detail.Surcharge.GetFloatAmount(),
			SurchargeCurrency: detail.Surcharge.GetCurrency(),
		},
		RateGuaranteed: detail.RateGuaranteed,
		ApproximateRate: &RateInfo{
			Base:              detail.ApproximateBase.GetFloatAmount(),
			BaseCurrency:      detail.ApproximateBase.GetCurrency(),
			Tax:               detail.ApproximateTax.GetFloatAmount(),
			TaxCurrency:       detail.ApproximateTax.GetCurrency(),
			Total:             detail.ApproximateTotal.GetFloatAmount(),
			TotalCurrency:     detail.ApproximateTotal.GetCurrency(),
			Surcharge:         detail.ApproximateSurcharge.GetFloatAmount(),
			SurchargeCurrency: detail.ApproximateSurcharge.GetCurrency(),
		},
		ApproximateRateGuaranteed: detail.ApproximateRateGuaranteed,
		RateCategory:              string(detail.RateCategory),
		Key:                       string(detail.Key),
		RateSupplier:              string(detail.RateSupplier),
		RateOfferId:               string(detail.RateOfferId),
		InPolicy:                  detail.InPolicy,
	}

	//解析酒店价格描述
	if detail.RoomRateDescription != nil && len(detail.RoomRateDescription) > 0 {
		result.Descriptions = make([]*RoomRateDescription, 0)
		for _, rrd := range detail.RoomRateDescription {
			desc := &RoomRateDescription{
				Name: rrd.Name,
				Text: rrd.Text,
			}
			result.Descriptions = append(result.Descriptions, desc)
		}
	}

	//解析日期价格
	if detail.HotelRateByDate != nil && len(detail.HotelRateByDate) > 0 {
		result.DateRates = make([]*HotelRateByDate, 0)
		for _, hrbd := range detail.HotelRateByDate {
			rate := translateHotelRateByDateResult(hrbd)
			if rate != nil {
				result.DateRates = append(result.DateRates, rate)
			}
		}
	}

	//解析折扣信息

	//解析税费详情
	if detail.TaxDetails != nil && detail.TaxDetails.Tax != nil &&
		len(detail.TaxDetails.Tax) > 0 {
		result.TaxDetails = make([]*TaxDetail, 0)

		for _, td := range detail.TaxDetails.Tax {
			tax := translateTaxDetailsResult(td)
			if tax != nil {
				result.TaxDetails = append(result.TaxDetails, tax)
			}
		}
	}

	//解析退订规定
	if detail.CancelInfo != nil {
		result.CancelRule = translateCancelInfoToCancelRule(detail.CancelInfo)
	}

	//解析担保信息
	if detail.GuaranteeInfo != nil {
		result.GuaranteeInfo = translateGuaranteeInfoResult(detail.GuaranteeInfo)
	}

	//解析匹配指示器
	if detail.RateMatchIndicator != nil && len(detail.RateMatchIndicator) > 0 {
		result.MatchInds = make([]*RateMatchIndicator, 0)
		for _, rmi := range detail.RateMatchIndicator {
			ind := &RateMatchIndicator{
				Type:   rmi.Type,
				Status: rmi.Status,
				Value:  rmi.Value,
			}
			result.MatchInds = append(result.MatchInds, ind)
		}
	}

	return result
}

//translateHotelRateByDateResult 转换HotelRateByDate
func translateHotelRateByDateResult(detail *hotrs.HotelRateByDate) *HotelRateByDate {
	if nil == detail {
		return nil
	}

	result := &HotelRateByDate{
		EffectiveDate: detail.EffectiveDate,
		ExpireDate:    detail.ExpireDate,
		Rate: &RateInfo{
			Base:              detail.Base.GetFloatAmount(),
			BaseCurrency:      detail.Base.GetCurrency(),
			Tax:               detail.Tax.GetFloatAmount(),
			TaxCurrency:       detail.Tax.GetCurrency(),
			Total:             detail.Total.GetFloatAmount(),
			TotalCurrency:     detail.Total.GetCurrency(),
			Surcharge:         detail.Surcharge.GetFloatAmount(),
			SurchargeCurrency: detail.Surcharge.GetCurrency(),
		},
		ApproximateRate: &RateInfo{
			Base:          detail.ApproximateBase.GetFloatAmount(),
			BaseCurrency:  detail.ApproximateBase.GetCurrency(),
			Total:         detail.ApproximateTotal.GetFloatAmount(),
			TotalCurrency: detail.ApproximateTotal.GetCurrency(),
		},
		Contents: detail.Contents,
	}

	return result
}

//translateTaxDetailsResult 转换税费详情
func translateTaxDetailsResult(tax *hotrs.Tax) *TaxDetail {
	if nil == tax {
		return nil
	}

	result := &TaxDetail{
		Amount:         tax.Amount.GetFloatAmount(),
		Currency:       tax.Amount.GetCurrency(),
		Percentage:     tax.Percentage,
		Code:           int(tax.Code),
		EffectiveDate:  tax.EffectiveDate,
		ExpireDate:     tax.ExpirationDate,
		Term:           tax.Term,
		CollectionFreq: tax.CollectionFreq,
	}

	return result
}

//translateCancelInfoToCancelRule 将CancelInfo转换为CancelRule
func translateCancelInfoToCancelRule(info *hotrs.CancelInfo) *CancelRule {
	if nil == info {
		return nil
	}

	result := &CancelRule{
		NonRefundable:                 string(info.NonRefundableStayIndicator),
		CancelDeadline:                info.CancelDeadline,
		TaxInclusive:                  info.TaxInclusive,
		FeeInclusive:                  info.FeeInclusive,
		CancelPenaltyAmount:           info.CancelPenaltyAmount.GetFloatAmount(),
		CancelPenaltyAmountCurrency:   info.CancelPenaltyAmount.GetCurrency(),
		NumberOfNights:                info.NumberOfNights,
		CancelPenaltyPercent:          info.CancelPenaltyPercent,
		CancelPenaltyPercentAppliesTo: info.CancelPenaltyPercentAppliesTo,
	}

	return result
}

//translateGuaranteeInfoResult 转换GuaranteeInfo
func translateGuaranteeInfoResult(info *hotrs.GuaranteeInfo) *GuaranteeInfo {
	if nil == info {
		return nil
	}

	result := &GuaranteeInfo{
		DepositNights:       info.DepositNights,
		DepositPercent:      info.DepositPercent,
		AbsoluteDeadline:    info.AbsoluteDeadline,
		CredentialsRequired: info.CredentialsRequired,
		HoldTime:            info.HoldTime,
		GuaranteeType:       info.GuaranteeType,
	}

	if info.DepositAmount != nil {
		result.DepositAmount = info.DepositAmount.Amount.GetFloatAmount()
		result.DepositAmountCurrency = info.DepositAmount.Amount.GetCurrency()
		result.ApproximateDepositAmount = info.DepositAmount.ApproximateAmount.GetFloatAmount()
		result.ApproximateDepositAmountCurrency = info.DepositAmount.ApproximateAmount.GetCurrency()
	}

	if info.GuaranteePaymentType != nil && len(info.GuaranteePaymentType) > 0 {
		result.GuaranteePaymentType = make([]*GuaranteePaymentType, 0)
		for _, gpt := range info.GuaranteePaymentType {
			guarantee := &GuaranteePaymentType{
				Value:       gpt.Value,
				Type:        gpt.Type,
				Description: gpt.Description,
			}
			result.GuaranteePaymentType = append(result.GuaranteePaymentType, guarantee)
		}
	}

	return result
}

//translateMediaItemResult ...
func translateMediaItemResult(item *comrs.MediaItem) *hotavmodel.MediaInfoResult {
	if nil == item {
		return nil
	}

	media := &hotavmodel.MediaInfoResult{
		Caption:  item.Caption,
		Height:   item.Height,
		Width:    item.Width,
		Type:     item.Type,
		Url:      item.Url,
		Icon:     item.Icon,
		SizeCode: item.SizeCode,
	}

	return media
}
