package interfaces

type Command interface {
	Execute()
	Undo()
	GetName() string
}
