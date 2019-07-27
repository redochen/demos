package enum

//Defines the ability to eticket an entity (Yes, No, Required, Ticketless)
type EnumMealService int

const (
	MsMeal EnumMealService = iota
	MsColdMeal
	MsHotMeal
	MsBreakfast
	MsContinentalBreakfast
	MsLunch
	MsDinner
	MsSnackOrBrunch
	MsFoodForPurchase
	MsComplimentaryRefreshments
	MsAlcoholicBeveragesForPurchase
	MsComplimentaryAlcoholicBeverages
	MsFoodAndBeveragesForPurchase
	MsNoMealService
	MsRefreshmentsForPurchase
)

//获取EnumMealService的字符串值
func (this EnumMealService) String() string {
	switch this {
	case MsMeal:
		return "Meal"
	case MsColdMeal:
		return "ColdMeal"
	case MsHotMeal:
		return "HotMeal"
	case MsBreakfast:
		return "Breakfast"
	case MsContinentalBreakfast:
		return "ContinentalBreakfast"
	case MsLunch:
		return "Lunch"
	case MsDinner:
		return "Dinner"
	case MsSnackOrBrunch:
		return "SnackOrBrunch"
	case MsFoodForPurchase:
		return "FoodForPurchase"
	case MsComplimentaryRefreshments:
		return "ComplimentaryRefreshments"
	case MsAlcoholicBeveragesForPurchase:
		return "AlcoholicBeveragesForPurchase"
	case MsComplimentaryAlcoholicBeverages:
		return "ComplimentaryAlcoholicBeverages"
	case MsFoodAndBeveragesForPurchase:
		return "FoodAndBeveragesForPurchase"
	case MsNoMealService:
		return "NoMealService"
	case MsRefreshmentsForPurchase:
		return "RefreshmentsForPurchase"
	}
	return ""
}

//获取EnumMealService的整数值
func (this EnumMealService) Value() int {
	if this >= MsMeal && this <= MsRefreshmentsForPurchase {
		return int(this)
	} else {
		return -1
	}
}
