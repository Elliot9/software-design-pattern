package core

import "fmt"

type Easel struct {
	Width      int
	Height     int
	Components []UIComponent
}

func NewEasel(width int, height int) *Easel {
	return &Easel{
		Width:      width,
		Height:     height,
		Components: []UIComponent{},
	}
}

func (e *Easel) AddComponent(component UIComponent) {
	e.Components = append(e.Components, component)
}

func (e *Easel) Render() {
	canvas := make([][]rune, e.Height)
	for i := range canvas {
		canvas[i] = make([]rune, e.Width)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}

	for _, component := range e.Components {
		componentLines := component.Render()
		x, y := component.GetPosition()
		for i, line := range componentLines {
			for j, char := range line {
				if y+i < e.Height && x+j < e.Width {
					canvas[y+i][x+j] = char
				}
			}
		}
	}

	for i := 0; i < e.Width; i++ {
		canvas[0][i] = '.'
		canvas[e.Height-1][i] = '.'
	}
	for i := 1; i < e.Height-1; i++ {
		canvas[i][0] = '.'
		canvas[i][e.Width-1] = '.'
	}

	for _, line := range canvas {
		fmt.Println(string(line))
	}
}
