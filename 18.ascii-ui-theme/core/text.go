package core

type Text interface {
	UIComponent
}

type BaseText struct {
	*BaseUIComponent
	Text string
}

func NewBaseText(text string, pointX int, pointY int) *BaseText {
	return &BaseText{
		BaseUIComponent: NewBaseUIComponent(pointX, pointY),
		Text:            text,
	}
}

var _ Text = &BaseText{}
