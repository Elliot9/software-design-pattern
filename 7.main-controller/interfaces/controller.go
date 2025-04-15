package interfaces

type Controller interface {
	GetKeyboard() map[string]Command
	SetKeyboard(map[string]Command)
	Reset()
}
