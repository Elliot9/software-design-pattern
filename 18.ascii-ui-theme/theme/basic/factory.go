package basic

import "github/elliot9/class18/core"

type BasicThemeFactory struct {
	core.ThemeFactory
}

func NewBasicThemeFactory() *BasicThemeFactory {
	return &BasicThemeFactory{}
}

func (f *BasicThemeFactory) CreateButton(text string, paddingWidth int, paddingHeight int, pointX int, pointY int) core.Button {
	return NewBasicButton(text, paddingWidth, paddingHeight, pointX, pointY)
}

func (f *BasicThemeFactory) CreateNumberedList(texts []string, pointX int, pointY int) core.NumberedList {
	return NewBasicNumberedList(texts, pointX, pointY)
}

func (f *BasicThemeFactory) CreateText(text string, pointX int, pointY int) core.Text {
	return NewBasicText(text, pointX, pointY)
}

var _ core.ThemeFactory = &BasicThemeFactory{}
