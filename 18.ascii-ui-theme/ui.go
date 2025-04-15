package main

import "github/elliot9/class18/core"

type AsciiUI struct {
	*core.Easel
	theme core.ThemeFactory
}

func NewAsciiUI(width int, height int) *AsciiUI {
	return &AsciiUI{
		Easel: core.NewEasel(width, height),
	}
}

func (ui *AsciiUI) SetTheme(theme core.ThemeFactory) {
	ui.theme = theme
}

func (ui *AsciiUI) AddButton(text string, paddingWidth int, paddingHeight int, pointX int, pointY int) {
	ui.AddComponent(ui.theme.CreateButton(text, paddingWidth, paddingHeight, pointX, pointY))
}

func (ui *AsciiUI) AddText(text string, pointX int, pointY int) {
	ui.AddComponent(ui.theme.CreateText(text, pointX, pointY))
}

func (ui *AsciiUI) AddNumberedList(texts []string, pointX int, pointY int) {
	ui.AddComponent(ui.theme.CreateNumberedList(texts, pointX, pointY))
}

func (ui *AsciiUI) Render() {
	ui.Easel.Render()
}
