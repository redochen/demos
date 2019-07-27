package air

import (
	. "github.com/redochen/demos/travelport-uapi/models/air"
	"github.com/redochen/demos/travelport-uapi/soap/air"
	airrq "github.com/redochen/demos/travelport-uapi/soap/air/request"
	airrs "github.com/redochen/demos/travelport-uapi/soap/air/response"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	. "github.com/redochen/tools/string"
	"strconv"
	"strings"
)

//GetAirSegmentParam 获取航段请求参数
func GetAirSegmentParam(segment *Segment, index int, providerCode string, specifyBookingCode bool, pricingCmd *airrq.AirPricingCommand) *airrq.AirSegment {
	if nil == segment {
		return nil
	}

	airSegment := &airrq.AirSegment{
		Group:                     segment.Group,
		Carrier:                   segment.Carrier,
		FlightNumber:              com.TypeFlightNumber(segment.FlightNumber),
		Origin:                    com.TypeIATACode(segment.DepartureAirport),
		Destination:               com.TypeIATACode(segment.ArrivalAirport),
		DepartureTime:             segment.DepartureTime,
		ArrivalTime:               segment.ArrivalTime,
		ClassOfService:            com.TypeClassOfService(segment.BookingCode),
		Seamless:                  segment.Seamless,
		LinkAvailability:          segment.LinkAvailability,
		AvailabilityDisplayType:   segment.AvailabilityDisplayType,
		AvailabilitySource:        air.TypeAvailabilitySource(segment.AvailabilitySource),
		OptionalServicesIndicator: segment.OptionalServicesIndicator,
		ParticipantLevel:          segment.ParticipantLevel,
		PolledAvailabilityOption:  segment.PolledAvailabilityOption,
	}

	//根据以下文档如果是Connection，则要创建Connection节点：
	//https://support.travelport.com/webhelp/uapi/uAPI.htm#Air/Shared_Air_Topics/AirSegmentConnectionLogic.htm
	if segment.IsConnection {
		airSegment.Connection = &airrq.Connection{}
	}

	if segment.MarriageGroup > 0 {
		airSegment.MarriageGroup = segment.MarriageGroup
	}

	if len(segment.ProviderCode) > 0 {
		airSegment.ProviderCode = com.TypeProviderCode(segment.ProviderCode)
	} else {
		//segment.ProviderCode = com.TypeProviderCode("1G")
		airSegment.ProviderCode = com.TypeProviderCode(providerCode)
	}

	airSegment.TravelOrder = index + 1

	if len(segment.SegmentKey) > 0 {
		airSegment.Key = com.TypeRef(segment.SegmentKey)
	} else {
		airSegment.Key = com.TypeRef(strconv.Itoa(index))
	}

	//指定要预订的舱位
	if specifyBookingCode && pricingCmd != nil && segment.BookingCode != "" {
		pricingCmd.SetClassOfService(airSegment.Key, segment.BookingCode)
	}

	return airSegment
}

//GetAirSearchModifiers 获取查询选项
func GetAirSearchModifiers(searchModifiers *SearchModifiers, providerCode string) *airrq.AirSearchModifiers {
	if nil == searchModifiers {
		return nil
	}

	modifiers := &airrq.AirSearchModifiers{
		PreferredProviders: &airrq.PreferredProviders{
			Provider: make([]*comrq.Provider, 0),
		},
	}

	//指定1G
	modifiers.PreferredProviders.Provider = append(modifiers.PreferredProviders.Provider,
		&comrq.Provider{
			Code: providerCode, //"1G",
		})

	//设置转机次数及航司信息
	sameAirline := (searchModifiers.SameAirline == 1)
	sameAirlineSpecified := (searchModifiers.SameAirline == 1 || searchModifiers.SameAirline == 2)
	setFlightType(modifiers, searchModifiers.NumOfStops, sameAirline, sameAirlineSpecified)

	//最大转机时间
	if searchModifiers.MaxConnectionMintues > 0 {
		modifiers.MaxConnectionTime = searchModifiers.MaxConnectionMintues
	}

	//最大旅行时间
	if searchModifiers.MaxJourneyHours > 0 {
		modifiers.MaxJourneyTime = uint(searchModifiers.MaxJourneyHours)
	}

	//设置包含的航司
	if searchModifiers.IncludeAirlines != nil && len(searchModifiers.IncludeAirlines) > 0 {
		modifiers.PermittedCarriers = &airrq.PermittedCarriers{
			Carrier: make([]*comrq.Carrier, 0),
		}
		for _, ia := range searchModifiers.IncludeAirlines {
			ic := &comrq.Carrier{
				Code: ia,
			}

			modifiers.PermittedCarriers.Carrier = append(
				modifiers.PermittedCarriers.Carrier,
				ic)
		}
	}

	//设置排除的航司
	if searchModifiers.ExcludeAirlines != nil && len(searchModifiers.ExcludeAirlines) > 0 {
		modifiers.ProhibitedCarriers = &airrq.ProhibitedCarriers{
			Carrier: make([]*comrq.Carrier, 0),
		}
		for _, ea := range searchModifiers.ExcludeAirlines {
			ec := &comrq.Carrier{
				Code: ea,
			}

			modifiers.ProhibitedCarriers.Carrier = append(
				modifiers.ProhibitedCarriers.Carrier,
				ec)
		}
	}

	return modifiers
}

//setFlightType 设置行程类型
func setFlightType(modifiers *airrq.AirSearchModifiers, numOfStops int, sameAirline, sameAirlineSpecified bool) {
	if nil == modifiers {
		return
	}

	//MaxConnections and MaxStops are not supported.
	modifiers.FlightType = &airrq.FlightType{
		MaxConnections: numOfStops,
		StopDirects:    true,
		NonStopDirects: true,
	}

	if !sameAirlineSpecified {
		return
	}

	modifiers.FlightType.RequireSingleCarrier = sameAirline

	if numOfStops <= 0 {
		return
	}

	if sameAirline {
		setOnlineFlightType(modifiers, numOfStops)
	} else {
		setInterlineFlightType(modifiers, numOfStops)
	}
}

//setOnlineFlightType 设置同一航司的行程类型
func setOnlineFlightType(modifiers *airrq.AirSearchModifiers, numOfStops int) {
	if nil == modifiers {
		return
	}

	switch numOfStops {
	case 1:
		modifiers.FlightType.SingleOnlineCon = true
		break
	case 2:
		modifiers.FlightType.DoubleOnlineCon = true
		break
	case 3:
		modifiers.FlightType.TripleOnlineCon = true
		break
	default:
		break
	}
}

//setInterlineFlightType 设置不同航司的行程类型
func setInterlineFlightType(modifiers *airrq.AirSearchModifiers, numOfStops int) {
	if nil == modifiers {
		return
	}

	switch numOfStops {
	case 1:
		modifiers.FlightType.SingleInterlineCon = true
		break
	case 2:
		modifiers.FlightType.DoubleInterlineCon = true
		break
	case 3:
		modifiers.FlightType.TripleInterlineCon = true
		break
	default:
		break
	}
}

//GetSegmentResult 获取航段结果
func GetSegmentResult(airSegment *airrs.AirSegment, index int, isConnection bool, optionKey string, flightDetails *airrs.FlightDetailsList) *Segment {
	if nil == airSegment {
		return nil
	}

	segment := &Segment{
		Group:                     airSegment.Group,
		Index:                     index,
		OptionKey:                 optionKey,
		SegmentKey:                string(airSegment.Key),
		Carrier:                   airSegment.Carrier,
		FlightNumber:              int(airSegment.FlightNumber),
		DepartureAirport:          string(airSegment.Origin),
		ArrivalAirport:            string(airSegment.Destination),
		TravelTime:                airSegment.TravelTime,
		DepartureTime:             airSegment.DepartureTime,
		ArrivalTime:               airSegment.ArrivalTime,
		FlightTime:                airSegment.FlightTime,
		Distance:                  airSegment.Distance,
		Equipment:                 string(airSegment.Equipment),
		ETicket:                   strings.EqualFold(strings.ToLower(airSegment.ETicketability), "yes"),
		ChangeOfPlane:             airSegment.ChangeOfPlane,
		CabinClass:                airSegment.CabinClass,
		BookingCode:               string(airSegment.ClassOfService),
		MarriageGroup:             airSegment.MarriageGroup,
		Status:                    airSegment.Status,
		ReservationInfoRef:        string(airSegment.ProviderReservationInfoRef),
		TravelOrder:               airSegment.TravelOrder,
		Seamless:                  airSegment.Seamless,
		IsConnection:              isConnection,
		LinkAvailability:          airSegment.LinkAvailability,
		AvailabilityDisplayType:   airSegment.AvailabilityDisplayType,
		AvailabilitySource:        string(airSegment.AvailabilitySource),
		OptionalServicesIndicator: airSegment.OptionalServicesIndicator,
		ParticipantLevel:          airSegment.ParticipantLevel,
		PolledAvailabilityOption:  airSegment.PolledAvailabilityOption,
		ProviderCode:              string(airSegment.ProviderCode),
	}

	//解析共享航班信息
	if airSegment.CodeshareInfo != nil {
		segment.OperatingCarrier = string(airSegment.CodeshareInfo.OperatingCarrier)
		segment.OperatingFlightNumber = int(airSegment.CodeshareInfo.OperatingFlightNumber)
	}

	//解析航班详情
	parseFlightDetails(segment, airSegment, flightDetails)

	//解析AV信息
	segment.AvailInfos = GetAvailInfos(airSegment.AirAvailInfo)

	return segment
}

//ParseFlightDetails 解析航班详情。
// LowFareSearch和AirAvail接口请传入flightDetails；
// 其他接口（AirPrice和CreatePnr等）不用传入flightDetails。
func parseFlightDetails(segment *Segment, airSegment *airrs.AirSegment, flightDetails *airrs.FlightDetailsList) {
	if nil == segment || nil == airSegment {
		return
	}

	detailsArray := airSegment.FlightDetails

	//FlightDetails列表直接在segment下（出现在AirPrice和CreatePnr接口）
	if nil == detailsArray || len(detailsArray) <= 0 {
		//FlightDetails有个单独的列表（出现在LowFareSearch和AirAvail接口）
		if nil == airSegment.FlightDetailsRef || len(airSegment.FlightDetailsRef) <= 0 {
			return
		}

		segment.FlightDetailsKeys = make([]string, 0)

		//解析航班详情键列表
		for _, fdr := range airSegment.FlightDetailsRef {
			if fdr != nil {
				segment.FlightDetailsKeys = append(segment.FlightDetailsKeys, string(fdr.Key))
			}
		}

		if nil == flightDetails {
			return
		}

		detailsArray = flightDetails.GetAllFlightDetails(airSegment.FlightDetailsRef)
	}

	if nil == detailsArray || len(detailsArray) <= 0 {
		return
	}

	if len(detailsArray) == 1 { //无经停点
		parseFlightDetailsWithoutStopover(segment, detailsArray[0])
	} else { //有经停点
		parseFlightDetailsWithStopovers(segment, detailsArray)
	}
}

//parseFlightDetailsWithoutStopover 解析非经停航班详情
func parseFlightDetailsWithoutStopover(segment *Segment, details *airrs.FlightDetails) {
	if nil == segment || nil == details {
		return
	}

	segment.DepartureTerminal = details.OriginTerminal
	segment.ArrivalTerminal = details.DestinationTerminal

	if len(segment.Equipment) <= 0 {
		segment.Equipment = string(details.Equipment)
	}

	if segment.FlightTime <= 0 {
		segment.FlightTime = details.FlightTime
	}

	if segment.TravelTime <= 0 {
		segment.TravelTime = details.TravelTime
	}

	if segment.Distance <= 0 {
		segment.Distance = details.Distance
	}
}

//parseFlightDetailsWithStopovers 解析有经停点的航班信息
func parseFlightDetailsWithStopovers(segment *Segment, detailsArray []*airrs.FlightDetails) {
	if nil == segment || nil == detailsArray || len(detailsArray) <= 0 {
		return
	}

	var lasArrivalAirport string //上一个到达机场
	var lastArrivalTime string   //上一个到达时间
	var lastEquipment string
	var firstDepartureTerminal string
	var lastArrivalTerminal string

	segment.Stopovers = make([]*Stopover, 0)

	for _, details := range detailsArray {
		if nil == details {
			continue
		}

		if len(lasArrivalAirport) > 0 {
			stopover := &Stopover{
				Airport:       lasArrivalAirport,
				Equipment:     lastEquipment,
				ArrivalTime:   lastArrivalTime,
				DepartureTime: details.DepartureTime,
			}

			if len(stopover.Equipment) <= 0 {
				stopover.Equipment = string(details.Equipment)
			}

			lastArrival := CcStr.ParseTime(lastArrivalTime, "yyyy-MM-ddTHH:mm:ss.fffzzz", true)
			curDeparture := CcStr.ParseTime(details.DepartureTime, "yyyy-MM-ddTHH:mm:ss.fffzzz", true)
			if !curDeparture.IsZero() && !lastArrival.IsZero() {
				stopover.Duration = int(curDeparture.Sub(lastArrival).Minutes())
			}

			segment.Stopovers = append(segment.Stopovers, stopover)
		}

		lasArrivalAirport = string(details.Destination)
		lastArrivalTime = details.ArrivalTime
		lastEquipment = string(details.Equipment)

		lastArrivalTerminal = details.DestinationTerminal
		if len(firstDepartureTerminal) <= 0 {
			firstDepartureTerminal = details.OriginTerminal
		}

		if len(segment.Equipment) <= 0 {
			segment.Equipment = string(details.Equipment)
		}
	}

	segment.DepartureTerminal = firstDepartureTerminal
	segment.ArrivalTerminal = lastArrivalTerminal
}

//GetAvailInfos 解析余座信息
func GetAvailInfos(info []*airrs.AirAvailInfo) []*AvailInfo {
	if nil == info || len(info) <= 0 {
		return nil
	}

	avails := make([]*AvailInfo, 0)

	for _, aai := range info {
		if nil == aai {
			continue
		}

		if nil == aai.BookingCodeInfo || len(aai.BookingCodeInfo) <= 0 {
			continue
		}

		avail := &AvailInfo{
			ProviderCode: string(aai.ProviderCode),
			AvailCabins:  make([]*AvailCabin, 0),
		}

		for _, bci := range aai.BookingCodeInfo {
			cabins := GetAvailCabins(bci)
			if nil == cabins || len(cabins) <= 0 {
				continue
			}

			for _, cabin := range cabins {
				avail.AvailCabins = append(avail.AvailCabins, cabin)
			}
		}

		//舱位都不可用时，直接返回nil
		if len(avail.AvailCabins) == 0 {
			return nil
		}

		avails = append(avails, avail)
	}

	return avails
}

//GetAvailCabins 解析余座数：F9|A9|C9|D9|Z0|Y9|B9|M9|H9|K9|L9|Q9|G9|V9|U9
func GetAvailCabins(info *airrs.BookingCodeInfo) []*AvailCabin {
	if nil == info || len(info.BookingCounts) == 0 {
		return nil
	}

	availCabins := strings.Split(info.BookingCounts, "|")
	if nil == availCabins || len(availCabins) == 0 {
		return nil
	}

	list := make([]*AvailCabin, 0)

	for _, ac := range availCabins {
		if len(ac) != 2 {
			continue
		}

		cabin := ac[0:1]
		count := CcStr.ParseInt(ac[1:2])
		ch := ac[1]

		if count > 0 || ch == 'a' || ch == 'A' {
			bci := &AvailCabin{
				CabinClass:  info.CabinClass,
				BookingCode: cabin,
				AvailCount:  string(ch),
			}

			list = append(list, bci)
		}
	}

	return list
}

//GetPricingInfo 转换AirPricingInfo
func GetPricingInfo(info *airrs.AirPricingInfo, fareInfos *airrs.FareInfoList, pointKey string) *PricingInfo {
	price := &PricingInfo{
		PricePointKey:  pointKey,
		PricingInfoKey: string(info.Key),
		Price: &Price{
			BasePrice:         info.BasePrice.GetFloatAmount(),
			BaseCurrency:      info.BasePrice.GetCurrency(),
			AppBasePrice:      info.ApproximateBasePrice.GetFloatAmount(),
			AppBaseCurrency:   info.ApproximateBasePrice.GetCurrency(),
			EquivBasePrice:    info.EquivalentBasePrice.GetFloatAmount(),
			EquivBaseCurrency: info.EquivalentBasePrice.GetCurrency(),
			Taxes:             info.Taxes.GetFloatAmount(),
			TaxesCurrency:     info.Taxes.GetCurrency(),
			AppTaxes:          info.ApproximateTaxes.GetFloatAmount(),
			AppTaxesCurrency:  info.ApproximateTaxes.GetCurrency(),
			TotalPrice:        info.TotalPrice.GetFloatAmount(),
			TotalCurrency:     info.TotalPrice.GetCurrency(),
			AppTotalPrice:     info.ApproximateTotalPrice.GetFloatAmount(),
			AppTotalCurrency:  info.ApproximateTotalPrice.GetCurrency(),
		},
		PlatingCarrier:      string(info.PlatingCarrier),
		LatestTicketingTime: info.LatestTicketingTime,
		FareCalc:            string(info.FareCalc),
		FareInfos:           parseFareInfo(info.FareInfo),
		TaxInfos:            parseTaxInfo(info.TaxInfo),
		PricingMethod:       info.PricingMethod,
		Exchangeable:        info.Exchangeable,
		Refundable:          info.Refundable,
		BaggageInfos:        parseBaggageInfo(info.BaggageAllowances),
		ProviderCode:        string(info.ProviderCode),
		IncludesVAT:         info.IncludesVAT,
	}

	//改期规定
	if info.ChangePenalty != nil {
		if info.ChangePenalty.Amount != "" {
			price.ChangePenaltyAmount = CcStr.FormatFloat(info.ChangePenalty.Amount.GetFloatAmount())
			price.ChangePenaltyCurrency = info.ChangePenalty.Amount.GetCurrency()
		} else {
			price.ChangePenaltyPercent = string(info.ChangePenalty.Percentage)
		}
	}

	//退票规定
	if info.CancelPenalty != nil {
		if info.CancelPenalty.Amount != "" {
			price.CancelPenaltyAmount = CcStr.FormatFloat(info.CancelPenalty.Amount.GetFloatAmount())
			price.CancelPenaltyCurrency = info.CancelPenalty.Amount.GetCurrency()
		} else {
			price.CancelPenaltyPercent = string(info.CancelPenalty.Percentage)
		}
	}

	//行李额规定
	if info.FareInfoRef != nil && len(info.FareInfoRef) > 0 {
		price.BaggageInfos = getBaggageInfo(fareInfos, info.FareInfoRef)
		price.PrivateFareType = fareInfos.GetPrivateFareType(info.FareInfoRef[0].Key)
	}

	//乘客类型
	if info.PassengerType != nil && len(info.PassengerType) > 0 {
		price.PassengerType = string(info.PassengerType[0].Code)
	}

	return price
}

//parseFareInfo 解析运价信息
func parseFareInfo(fareInfos []*airrs.FareInfo) []*FareInfo {
	if nil == fareInfos || len(fareInfos) <= 0 {
		return nil
	}

	result := make([]*FareInfo, 0)
	for _, fareInfo := range fareInfos {
		info := &FareInfo{
			FareInfoKey:     string(fareInfo.Key),
			FareBasis:       fareInfo.FareBasis,
			PassengerType:   string(fareInfo.PassengerTypeCode),
			Amount:          fareInfo.Amount.GetFloatAmount(),
			Currency:        fareInfo.Amount.GetCurrency(),
			Origin:          string(fareInfo.Origin),
			Destination:     string(fareInfo.Destination),
			DepartureDate:   fareInfo.DepartureDate,
			NotValidBefore:  fareInfo.NotValidBefore,
			NotValidAfter:   fareInfo.NotValidAfter,
			EffectiveDate:   fareInfo.EffectiveDate,
			PrivateFareType: fareInfo.PrivateFare,
		}

		result = append(result, info)
	}

	return result
}

//parseTaxInfo 解析税收信息
func parseTaxInfo(taxInfos []*airrs.TaxInfo) []*TaxInfo {
	if nil == taxInfos || len(taxInfos) <= 0 {
		return nil
	}

	result := make([]*TaxInfo, 0)
	for _, taxInfo := range taxInfos {
		info := &TaxInfo{
			TaxInfoKey: string(taxInfo.Key),
			Category:   taxInfo.Category,
			Amount:     taxInfo.Amount.GetFloatAmount(),
			Currency:   taxInfo.Amount.GetCurrency(),
		}

		result = append(result, info)
	}

	return result
}

//getBaggageInfo 获取行李额规定
func getBaggageInfo(fareInfos *airrs.FareInfoList, fareInfoRefs []*airrs.FareInfoRef) []*BaggageInfo {
	if nil == fareInfos {
		return nil
	}

	if nil == fareInfoRefs || len(fareInfoRefs) == 0 {
		return nil
	}

	baggages := make([]*BaggageInfo, 0)

	for _, fareRef := range fareInfoRefs {
		fare := fareInfos.GetFareInfo(fareRef.Key)
		if nil == fare || nil == fare.BaggageAllowance {
			continue
		}

		baggage := &BaggageInfo{
			PassengerType: string(fare.PassengerTypeCode),
			Origin:        string(fare.Origin),
			Destination:   string(fare.Destination),
		}

		if len(fare.BaggageAllowance.NumberOfPieces) > 0 {
			baggage.NumberOfPieces = CcStr.ParseInt(fare.BaggageAllowance.NumberOfPieces)
		}

		if fare.BaggageAllowance.MaxWeight != nil && fare.BaggageAllowance.MaxWeight.Value > 0 {
			baggage.MaxWeight = &BaggageMeasure{
				Type:  "Weight",
				Value: float32(fare.BaggageAllowance.MaxWeight.Value),
				Unit:  fare.BaggageAllowance.MaxWeight.Unit,
			}
		}

		if baggage != nil {
			baggages = append(baggages, baggage)
		}
	}

	return baggages
}

//parseBaggageInfo 解析行李规定
func parseBaggageInfo(baggageInfo *airrs.BaggageAllowances) []*BaggageInfo {
	if nil == baggageInfo {
		return nil
	}

	result := make([]*BaggageInfo, 0)

	allowance := parseBaggageAllowanceInfo(baggageInfo.BaggageAllowanceInfo)
	if allowance != nil && len(allowance) > 0 {
		for _, a := range allowance {
			result = append(result, a)
		}
	}

	carryOn := parseCarryOnAllowanceInfo(baggageInfo.CarryOnAllowanceInfo)
	if carryOn != nil && len(carryOn) > 0 {
		for _, c := range carryOn {
			result = append(result, c)
		}
	}

	return result
}

//parseBaggageAllowanceInfo 解析BaggageAllowanceInfo
func parseBaggageAllowanceInfo(baggageAllowanceInfo []*airrs.BaggageAllowanceInfo) []*BaggageInfo {
	if nil == baggageAllowanceInfo || len(baggageAllowanceInfo) <= 0 {
		return nil
	}

	result := make([]*BaggageInfo, 0)

	for _, info := range baggageAllowanceInfo {
		if nil == info {
			continue
		}

		baggage := &BaggageInfo{
			IsCarryOn:     false,
			PassengerType: string(info.TravelerType),
			Origin:        string(info.Origin),
			Destination:   string(info.Destination),
			Carrier:       string(info.Carrier),
		}

		//解析文本信息
		baggage.Text = getTextInfo(info.TextInfo)

		if info.BagDetails != nil && len(info.BagDetails) > 0 {
			baggage.Details = make([]*BaggageDetails, 0)
			for _, bagDetails := range info.BagDetails {
				if nil == bagDetails {
					continue
				}

				details := &BaggageDetails{
					ApplicableInfo: bagDetails.ApplicableBags,
					Restriction:    parseBaggageRestrictions(bagDetails.BaggageRestriction),
					Price: &Price{
						BasePrice:        bagDetails.BasePrice.GetFloatAmount(),
						BaseCurrency:     bagDetails.BasePrice.GetCurrency(),
						AppBasePrice:     bagDetails.ApproximateBasePrice.GetFloatAmount(),
						AppBaseCurrency:  bagDetails.ApproximateBasePrice.GetCurrency(),
						Taxes:            bagDetails.Taxes.GetFloatAmount(),
						TaxesCurrency:    bagDetails.Taxes.GetCurrency(),
						TotalPrice:       bagDetails.TotalPrice.GetFloatAmount(),
						TotalCurrency:    bagDetails.TotalPrice.GetCurrency(),
						AppTotalPrice:    bagDetails.ApproximateTotalPrice.GetFloatAmount(),
						AppTotalCurrency: bagDetails.ApproximateTotalPrice.GetCurrency(),
					},
				}

				baggage.Details = append(baggage.Details, details)
			}
		}

		result = append(result, baggage)
	}

	return result
}

//parseCarryOnAllowanceInfo 解析CarryOnAllowanceInfo
func parseCarryOnAllowanceInfo(carryOnAllowanceInfo []*airrs.CarryOnAllowanceInfo) []*BaggageInfo {
	if nil == carryOnAllowanceInfo || len(carryOnAllowanceInfo) <= 0 {
		return nil
	}

	result := make([]*BaggageInfo, 0)

	for _, info := range carryOnAllowanceInfo {
		if nil == info {
			continue
		}

		baggage := &BaggageInfo{
			IsCarryOn:   true,
			Origin:      string(info.Origin),
			Destination: string(info.Destination),
			Carrier:     string(info.Carrier),
		}

		//解析文本信息
		baggage.Text = getTextInfo(info.TextInfo)

		if info.CarryOnDetails != nil && len(info.CarryOnDetails) > 0 {
			baggage.Details = make([]*BaggageDetails, 0)
			for _, carryOnDetails := range info.CarryOnDetails {
				if nil == carryOnDetails {
					continue
				}

				details := &BaggageDetails{
					ApplicableInfo: carryOnDetails.ApplicableCarryOnBags,
					Restriction:    parseBaggageRestrictions(carryOnDetails.BaggageRestriction),
					Price: &Price{
						BasePrice:        carryOnDetails.BasePrice.GetFloatAmount(),
						BaseCurrency:     carryOnDetails.BasePrice.GetCurrency(),
						AppBasePrice:     carryOnDetails.ApproximateBasePrice.GetFloatAmount(),
						AppBaseCurrency:  carryOnDetails.ApproximateBasePrice.GetCurrency(),
						Taxes:            carryOnDetails.Taxes.GetFloatAmount(),
						TaxesCurrency:    carryOnDetails.Taxes.GetCurrency(),
						TotalPrice:       carryOnDetails.TotalPrice.GetFloatAmount(),
						TotalCurrency:    carryOnDetails.TotalPrice.GetCurrency(),
						AppTotalPrice:    carryOnDetails.ApproximateTotalPrice.GetFloatAmount(),
						AppTotalCurrency: carryOnDetails.ApproximateTotalPrice.GetCurrency(),
					},
				}

				baggage.Details = append(baggage.Details, details)
			}
		}

		result = append(result, baggage)
	}

	return result
}

//parseBaggageRestrictions 解析行李限制
func parseBaggageRestrictions(baggageRestrictions []*airrs.BaggageRestriction) []*BaggageRestriction {
	if nil == baggageRestrictions || len(baggageRestrictions) <= 0 {
		return nil
	}

	result := make([]*BaggageRestriction, 0)

	for _, br := range baggageRestrictions {
		if nil == br {
			continue
		}

		restric := &BaggageRestriction{}

		if br.MaxWeight != nil {
			restric.MaxWeight = &BaggageMeasure{
				Type:  "Weight",
				Value: br.MaxWeight.Value,
				Unit:  br.MaxWeight.Unit,
			}
		}

		if br.Dimension != nil && len(br.Dimension) > 0 {
			restric.Dimensions = make([]*BaggageMeasure, 0)

			for _, d := range br.Dimension {
				measure := &BaggageMeasure{
					Type:  d.Type,
					Value: d.Value,
					Unit:  d.Unit,
				}
				restric.Dimensions = append(restric.Dimensions, measure)
			}
		}

		//解析文本信息
		restric.Text = getTextInfo(br.TextInfo)

		result = append(result, restric)
	}

	return result
}

//getTextInfo 获取文本信息
func getTextInfo(textInfo []*airrs.TextInfo) []string {
	if nil == textInfo || len(textInfo) <= 0 {
		return nil
	}

	result := make([]string, 0)

	for _, ti := range textInfo {
		if nil == ti || nil == ti.Text {
			continue
		}

		for _, text := range ti.Text {
			result = append(result, string(text))
		}
	}

	return result
}
