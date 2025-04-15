package core

type NumberedList interface {
	UIComponent
}

type BaseNumberedList struct {
	*BaseUIComponent
	Texts []string
}

func NewBaseNumberedList(texts []string, pointX int, pointY int) *BaseNumberedList {
	return &BaseNumberedList{
		BaseUIComponent: NewBaseUIComponent(pointX, pointY),
		Texts:           texts,
	}
}

var _ NumberedList = &BaseNumberedList{}
