package card

import "fmt"

type UnoCard struct {
	number int
	color  Color
}

func NewUnoCard(number int, color Color) *UnoCard {
	return &UnoCard{number: number, color: color}
}

func (u *UnoCard) Identify() string {
	return fmt.Sprintf("%d-%s", u.number, u.color)
}

type Color string

const (
	Red    Color = "Red"
	Yellow Color = "Yellow"
	Green  Color = "Green"
	Blue   Color = "Blue"
)
