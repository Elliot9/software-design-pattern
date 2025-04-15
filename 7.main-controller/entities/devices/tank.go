package devices

import "fmt"

type Tank struct{}

func (t *Tank) MoveForward() {
	fmt.Println("The tank has moved forward.")
}

func (t *Tank) MoveBackward() {
	fmt.Println("The tank has moved backward.")
}

func NewTank() *Tank {
	return &Tank{}
}
