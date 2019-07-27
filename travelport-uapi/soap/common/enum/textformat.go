package enum

//Indicates the format of text used in the description e.g. unformatted  or html.
type EnumTextFormat int

const (
	PlainText EnumTextFormat = iota //Textual data that is in ASCII format.
	HTML                            //HTML formatted text.
)

//获取EnumTextFormat的字符串值
func (this EnumTextFormat) String() string {
	switch this {
	case PlainText:
		return "PlainText"
	case HTML:
		return "HTML"
	}
	return ""
}

//获取EnumTextFormat的整数值
func (this EnumTextFormat) Value() int {
	if this >= PlainText && this <= HTML {
		return int(this)
	} else {
		return -1
	}
}
