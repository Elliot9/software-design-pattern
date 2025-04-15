package core

type UIComponent interface {
	Render() [][]rune
	GetPosition() (int, int)
}

type BaseUIComponent struct {
	PointX int
	PointY int
}

func NewBaseUIComponent(pointX int, pointY int) *BaseUIComponent {
	return &BaseUIComponent{
		PointX: pointX,
		PointY: pointY,
	}
}

func (c *BaseUIComponent) GetPosition() (int, int) {
	return c.PointX, c.PointY
}

func (c *BaseUIComponent) Render() [][]rune {
	panic("not implemented")
}

var _ UIComponent = &BaseUIComponent{}
