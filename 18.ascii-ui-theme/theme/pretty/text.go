package pretty

import (
	"github/elliot9/class18/core"
	"strings"
)

type PrettyText struct {
	core.BaseText
}

func NewPrettyText(text string, pointX int, pointY int) *PrettyText {
	return &PrettyText{
		BaseText: core.BaseText{
			BaseUIComponent: core.NewBaseUIComponent(pointX, pointY),
			Text:            text,
		},
	}
}

func (t *PrettyText) Render() [][]rune {
	textlines := strings.Split(t.Text, "\n")
	lines := make([][]rune, len(textlines))
	for i, line := range textlines {
		lines[i] = []rune(strings.ToUpper(line))
	}
	return lines
}

var _ core.Text = &PrettyText{}
var _ core.UIComponent = &PrettyText{}
