package main

// 做出遊戲決策
type Player interface {
	GetHands() []Card
	GetName() string
	GetScore() int
	SetUsedPermission()
	UsedPermission() bool
	AddHands(card Card)
	AddScore()
	ReplaceHands(cards []Card)
	Show() *Card
	DecideToExchange() bool
	UseExchange(players []Player) (Player, error)
	NameSelf()
}

type BasePlayer struct {
	score      int
	hands      []Card
	order      int
	name       string
	permission bool
}

func (p *BasePlayer) GetHands() []Card {
	return p.hands
}

func (p *BasePlayer) setName(name string) {
	p.name = name
}

func (p *BasePlayer) GetName() string {
	return p.name
}

func (p *BasePlayer) GetScore() int {
	return p.score
}

func (p *BasePlayer) SetUsedPermission() {
	p.permission = true
}

func (p *BasePlayer) UsedPermission() bool {
	return p.permission
}

func (p *BasePlayer) AddHands(card Card) {
	p.hands = append(p.hands, card)
}

func (p *BasePlayer) AddScore() {
	p.score++
}

func (p *BasePlayer) ReplaceHands(cards []Card) {
	p.hands = cards
}
