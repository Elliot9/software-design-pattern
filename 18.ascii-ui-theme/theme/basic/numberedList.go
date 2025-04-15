package basic

import (
	"fmt"
	"github/elliot9/class18/core"
)

type BasicNumberedList struct {
	core.BaseNumberedList
}

func NewBasicNumberedList(texts []string, pointX int, pointY int) *BasicNumberedList {
	return &BasicNumberedList{
		BaseNumberedList: core.BaseNumberedList{
			BaseUIComponent: core.NewBaseUIComponent(pointX, pointY),
			Texts:           texts,
		},
	}
}

func (b *BasicNumberedList) Render() [][]rune {
	lines := make([][]rune, len(b.Texts))
	for i, text := range b.Texts {
		line := fmt.Sprintf("%d. %s", i+1, text)
		lines[i] = []rune(line)
	}
	return lines
}

var _ core.UIComponent = &BasicNumberedList{}
var _ core.NumberedList = &BasicNumberedList{}
