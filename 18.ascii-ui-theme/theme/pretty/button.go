package pretty

import (
	"github/elliot9/class18/core"
)

type PrettyButton struct {
	core.BaseButton
}

func NewPrettyButton(text string, paddingWidth int, paddingHeight int, pointX int, pointY int) *PrettyButton {
	return &PrettyButton{
		BaseButton: core.BaseButton{
			BaseUIComponent: core.NewBaseUIComponent(pointX, pointY),
			Text:            text,
			PaddingWidth:    paddingWidth,
			PaddingHeight:   paddingHeight,
		},
	}
}

func (b *PrettyButton) Render() [][]rune {
	width := len(b.Text) + 2*b.PaddingWidth + 2
	height := 3 + 2*b.PaddingHeight

	lines := make([][]rune, height)
	for i := range lines {
		lines[i] = make([]rune, width)
		if i == 0 {
			lines[i][0] = '┌'
			lines[i][width-1] = '┐'
			for j := 1; j < width-1; j++ {
				lines[i][j] = '─'
			}
		} else if i == height-1 {
			lines[i][0] = '└'
			lines[i][width-1] = '┘'
			for j := 1; j < width-1; j++ {
				lines[i][j] = '─'
			}
		} else {
			lines[i][0] = '│'
			lines[i][width-1] = '│'
			for j := 1; j < width-1; j++ {
				lines[i][j] = ' '
			}
			if i == height/2 {
				copy(lines[i][1+b.PaddingWidth:], []rune(b.Text))
			}
		}
	}
	return lines
}

var _ core.Button = &PrettyButton{}
var _ core.UIComponent = &PrettyButton{}
