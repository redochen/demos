package response

//The base type for all responses.
type BaseRsp struct {
	ResponseMessage []*ResponseMessage `xml:"ResponseMessage,omitempty"`     //minOccurs="0" maxOccurs="unbounded"
	TraceId         string             `xml:"TraceId,attr,omitempty"`        //use="optional" //Unique identifier for this atomic transaction traced by the user. Use is optional.
	TransactionId   string             `xml:"TransactionId,attr,omitempty"`  //use="optional" //System generated unique identifier for this atomic transaction.
	ResponseTime    int                `xml:"ResponseTime,attr,omitempty"`   //use="optional"  //The time (in ms) the system spent processing this request, not including transmission times.
	CommandHistory  string             `xml:"CommandHistory,attr,omitempty"` //use="optional" //HTTP link to download command history and debugging information of the request that generated this response. Must be enabled on the system.
}

type BaseSearchRsp struct {
	BaseRsp
	NextResultReference []*NextResultReference `xml:"NextResultReference,omitempty"` //minOccurs="0" maxOccurs="unbounded"
}
