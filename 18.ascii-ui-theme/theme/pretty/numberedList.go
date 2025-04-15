package pretty

import (
	"fmt"
	"github/elliot9/class18/core"
)

type PrettyNumberedList struct {
	core.BaseNumberedList
}

func NewPrettyNumberedList(texts []string, pointX int, pointY int) *PrettyNumberedList {
	return &PrettyNumberedList{
		BaseNumberedList: core.BaseNumberedList{
			BaseUIComponent: core.NewBaseUIComponent(pointX, pointY),
			Texts:           texts,
		},
	}
}

func (l *PrettyNumberedList) Render() [][]rune {
	romannumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	lines := make([][]rune, len(l.Texts))
	for i, text := range l.Texts {
		if i < len(romannumerals) {
			line := fmt.Sprintf("%s. %s", romannumerals[i], text)
			lines[i] = []rune(line)
		}
	}
	return lines
}

var _ core.NumberedList = &PrettyNumberedList{}
var _ core.UIComponent = &PrettyNumberedList{}
