package pretty

import (
	"github/elliot9/class18/core"
)

type PrettyThemeFactory struct {
	core.ThemeFactory
}

func NewPrettyThemeFactory() *PrettyThemeFactory {
	return &PrettyThemeFactory{}
}

func (f *PrettyThemeFactory) CreateButton(text string, paddingWidth int, paddingHeight int, pointX int, pointY int) core.Button {
	return NewPrettyButton(text, paddingWidth, paddingHeight, pointX, pointY)
}

func (f *PrettyThemeFactory) CreateNumberedList(texts []string, pointX int, pointY int) core.NumberedList {
	return NewPrettyNumberedList(texts, pointX, pointY)
}

func (f *PrettyThemeFactory) CreateText(text string, pointX int, pointY int) core.Text {
	return NewPrettyText(text, pointX, pointY)
}

var _ core.ThemeFactory = &PrettyThemeFactory{}
