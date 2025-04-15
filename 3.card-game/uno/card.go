package main

import "fmt"

// 表示一張撲克牌，供階級花色比較
type Card struct {
	number int
	color  Color
}

func (c *Card) getNumber() int {
	return c.number
}

func (c *Card) getColor() Color {
	return c.color
}
func (c *Card) String() string {
	return fmt.Sprintf("%s %d", c.color, c.number)
}

type Color string

const (
	Red    Color = "Red"
	Yellow Color = "Yellow"
	Green  Color = "Green"
	Blue   Color = "Blue"
)
