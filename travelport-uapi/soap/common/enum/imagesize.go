package enum

//The image size
type EnumImageSize int

const (
	ImageSizeT EnumImageSize = iota //Thumbnail
	ImageSizeI                      //Minimum
	ImageSizeS                      //Small
	ImageSizeM                      //Medium
	ImageSizeL                      //Large
	ImageSizeE                      //Extra Large
	ImageSizeG                      //Guaranteed
	ImageSizeF                      //Forced
	ImageSizeB                      //Big
	ImageSizeJ                      //Jumbo
	ImageSizeO                      //Original
	ImageSizeH                      //Huge
	ImageSizeC                      //Colossal
)

//获取EnumImageSize的字符串值
func (this EnumImageSize) String() string {
	switch this {
	case ImageSizeT:
		return "T"
	case ImageSizeI:
		return "I"
	case ImageSizeS:
		return "S"
	case ImageSizeM:
		return "M"
	case ImageSizeL:
		return "L"
	case ImageSizeE:
		return "E"
	case ImageSizeG:
		return "G"
	case ImageSizeF:
		return "F"
	case ImageSizeB:
		return "B"
	case ImageSizeJ:
		return "J"
	case ImageSizeO:
		return "O"
	case ImageSizeH:
		return "H"
	case ImageSizeC:
		return "C"
	}
	return ""
}

//获取EnumImageSize的整数值
func (this EnumImageSize) Value() int {
	if this >= ImageSizeT && this <= ImageSizeC {
		return int(this)
	} else {
		return -1
	}
}
