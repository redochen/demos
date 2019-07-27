package soap

import "reflect"

const (
	// NSCom com命名空间
	NSCom string = "http://www.travelport.com/schema/common_v40_0"

	//NSUniv univ命名空间
	NSUniv string = "http://www.travelport.com/schema/universal_v40_0"

	//NSSys sys命名空间
	NSSys string = "http://www.travelport.com/schema/system_v32_0"

	// NSAir air命名空间
	NSAir string = "http://www.travelport.com/schema/air_v40_0"

	//NSHotel hot命名空间
	NSHotel string = "http://www.travelport.com/schema/hotel_v40_0"
)

//INameSpace 命名空间接口
type INameSpace interface {
	SetComNameSpace()
	SetSysNameSpace()
	SetUnivNameSpace()
	SetAirNameSpace()
	SetHotNameSpace()
}

//NameSpace c命名空间类
type NameSpace struct {
	ComNamespace  string `xml:"xmlns:com,attr,omitempty"`
	SysNamespace  string `xml:"xmlns:sys,attr,omitempty"`
	UnivNamespace string `xml:"xmlns:univ,attr,omitempty"`
	AirNamespace  string `xml:"xmlns:air,attr,omitempty"`
	HotNamespace  string `xml:"xmlns:hot,attr,omitempty"`
}

//SetComNameSpace 设置com命名空间
func (nameSpace *NameSpace) SetComNameSpace() {
	nameSpace.ComNamespace = NSCom
}

//SetSysNameSpace 设置sys命名空间
func (nameSpace *NameSpace) SetSysNameSpace() {
	nameSpace.SysNamespace = NSSys
}

//SetUnivNameSpace 设置univ命名空间
func (nameSpace *NameSpace) SetUnivNameSpace() {
	nameSpace.UnivNamespace = NSUniv
}

//SetAirNameSpace 设置air命名空间
func (nameSpace *NameSpace) SetAirNameSpace() {
	nameSpace.AirNamespace = NSAir
}

//SetHotNameSpace 设置hot命名空间
func (nameSpace *NameSpace) SetHotNameSpace() {
	nameSpace.HotNamespace = NSHotel
}

//IHasMembersWithNameSpace 具有实现了NameSpace接口的成员
type IHasMembersWithNameSpace interface {
	SetAllNameSpaces(ns string)
}

//SetNameSpace 设置命名空间
func SetNameSpace(item interface{}, ns string) {
	if nil == item || reflect.ValueOf(item).IsNil() {
		return
	}

	if one, ok := item.(INameSpace); ok {
		switch ns {
		case "com":
			one.SetComNameSpace()
			break
		case "sys":
			one.SetSysNameSpace()
			break
		case "univ":
			one.SetUnivNameSpace()
			break
		case "air":
			one.SetAirNameSpace()
			break
		case "hot":
			one.SetHotNameSpace()
			break
		default:
			break
		}
	}
}

//SetNameSpaces 设置所有成员的命名空间
func SetNameSpaces(item interface{}, ns string) {
	if nil == item {
		return
	}

	value := reflect.ValueOf(item)
	if value.IsNil() {
		return
	}

	kind := value.Kind()

	//切片或数组
	if reflect.Slice == kind || reflect.Array == kind {
		len := value.Len()
		if len <= 0 {
			return
		}

		for i := 0; i < len; i++ {
			v := value.Index(i).Interface()
			SetNameSpaces(v, ns)
		}
	} else if all, ok := item.(IHasMembersWithNameSpace); ok {
		all.SetAllNameSpaces(ns)
	} else {
		SetNameSpace(item, ns)
	}
}
