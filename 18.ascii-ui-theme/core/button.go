package core

type Button interface {
	UIComponent
}

type BaseButton struct {
	*BaseUIComponent
	Text          string
	PaddingWidth  int
	PaddingHeight int
}

func NewBaseButton(text string, paddingWidth int, paddingHeight int, pointX int, pointY int) *BaseButton {
	return &BaseButton{
		BaseUIComponent: NewBaseUIComponent(pointX, pointY),
		Text:            text,
		PaddingWidth:    paddingWidth,
		PaddingHeight:   paddingHeight,
	}
}

var _ Button = &BaseButton{}
