package basic

import (
	"github/elliot9/class18/core"
	"strings"
)

type BasicText struct {
	core.BaseText
}

func NewBasicText(text string, pointX int, pointY int) *BasicText {
	return &BasicText{
		BaseText: core.BaseText{
			BaseUIComponent: core.NewBaseUIComponent(pointX, pointY),
			Text:            text,
		},
	}
}

func (t *BasicText) Render() [][]rune {
	textlines := strings.Split(t.BaseText.Text, "\n")
	lines := make([][]rune, len(textlines))
	for i, line := range textlines {
		lines[i] = []rune(line)
	}
	return lines
}

var _ core.UIComponent = &BasicText{}
var _ core.Text = &BasicText{}
