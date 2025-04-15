package core

type ThemeFactory interface {
	CreateButton(text string, paddingWidth int, paddingHeight int, pointX int, pointY int) Button
	CreateNumberedList(texts []string, pointX int, pointY int) NumberedList
	CreateText(text string, pointX int, pointY int) Text
}
