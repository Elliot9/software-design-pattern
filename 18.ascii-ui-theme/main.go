package main

import (
	"github/elliot9/class18/theme/basic"
)

func main() {
	ui := NewAsciiUI(22, 13)
	ui.SetTheme(basic.NewBasicThemeFactory())
	// ui.SetTheme(pretty.NewPrettyThemeFactory())
	ui.AddButton("Hi, I miss u", 1, 0, 3, 1)
	ui.AddText("Do u love me ?\nPlease tell...", 4, 4)
	ui.AddButton("No", 1, 0, 3, 6)
	ui.AddButton("Yes", 1, 0, 12, 6)
	ui.AddNumberedList([]string{
		"Let's travel",
		"Back to home",
		"Have dinner",
	}, 3, 9)

	ui.Render()
}
