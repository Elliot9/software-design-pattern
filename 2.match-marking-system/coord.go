package main

type Coord struct {
	x int
	y int
}

func (c *Coord) GetX() int {
	return c.x
}

func (c *Coord) GetY() int {
	return c.y
}
