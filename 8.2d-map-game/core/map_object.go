package core

type MapObject interface {
	SetMap(mapper *Map)
	GetMap() *Map
	GetSymbol() string
	GetPosition() (int, int)
}

type BaseMapObject struct {
	symbol string
	mapper *Map
	MapObject
}

func (m *BaseMapObject) GetSymbol() string {
	return m.symbol
}

func (m *BaseMapObject) SetMap(mapper *Map) {
	m.mapper = mapper
}

func (m *BaseMapObject) GetMap() *Map {
	return m.mapper
}

func (m *BaseMapObject) GetPosition() (int, int) {
	for y, row := range m.GetMap().grid {
		for x, obj := range row {
			if obj == m.MapObject {
				return x, y
			}
		}
	}
	return -1, -1
}

type Obstacle struct {
	BaseMapObject
}

func NewObstacle() *Obstacle {
	o := &Obstacle{
		BaseMapObject: BaseMapObject{
			symbol: "□",
		},
	}
	o.MapObject = o
	return o
}

type Treasure struct {
	BaseMapObject
	Name        string
	Probability float64
	State       State
}

func NewTreasure(name string, probability float64, state State) *Treasure {
	t := &Treasure{
		BaseMapObject: BaseMapObject{
			symbol: "x",
		},
		Name:        name,
		Probability: probability,
		State:       state,
	}
	t.MapObject = t
	return t
}

var _ MapObject = &BaseMapObject{}
var _ MapObject = &Obstacle{}
var _ MapObject = &Treasure{}
