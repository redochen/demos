package air

//Stopover 经停类
type Stopover struct {
	Airport       string `xml:"airport" json:"airport"`                       //经停机场
	Duration      int    `xml:"duration,omitempty" json:"duration,omitempty"` //停留时间
	Equipment     string `xml:"equipment" json:"equipment"`                   //机型
	ArrivalTime   string `xml:"arrTime" json:"arrTime"`                       //降落时间，格式：yyyy-MM-ddTHH:mm:ss.fffzzz
	DepartureTime string `xml:"depTime" json:"depTime"`                       //起飞时间，格式：yyyy-MM-ddTHH:mm:ss.fffzzz
}
