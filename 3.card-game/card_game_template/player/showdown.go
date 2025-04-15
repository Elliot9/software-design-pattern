package player

import (
	"fmt"
	"github/elliot9/card_game_template/card"
	"math/rand"
)

type ShowdownPlayer interface {
	Player
	AddScore()
	GetScore() int
	UsedPermission() bool
	SetUsedPermission()
	ReplaceHands(cards []card.CardStrategy)
	DecideToExchange() bool
	UseExchange(players []ShowdownPlayer) (ShowdownPlayer, error)
}

type BaseShowdownPlayer struct {
	BasePlayer
	score      int
	permission bool
}

func (p *BaseShowdownPlayer) AddScore() {
	p.score++
}

func (p *BaseShowdownPlayer) GetScore() int {
	return p.score
}

func (p *BaseShowdownPlayer) UsedPermission() bool {
	return p.permission
}

func (p *BaseShowdownPlayer) SetUsedPermission() {
	p.permission = true
}

func (p *BaseShowdownPlayer) ReplaceHands(cards []card.CardStrategy) {
	p.hands = cards
}

type HumanShowdownPlayer struct {
	BaseShowdownPlayer
	Cli CLI
}

func (p *HumanShowdownPlayer) NameSelf() {
	fmt.Println("請輸入有效的名子:")
	name, _ := p.Cli.ReadLine()
	p.setName(name)
}

func (p *HumanShowdownPlayer) Show() card.CardStrategy {
	if len(p.GetHands()) == 0 {
		fmt.Printf("%s 沒有卡牌可以出\n", p.GetName())
		return nil
	}

	fmt.Println("輪到你出牌")
	for i, card := range p.GetHands() {
		fmt.Printf("%d: %s \n", i+1, card.Identify())
	}
	fmt.Println("請輸入要出牌的卡牌編號:")

	index, _ := p.Cli.ReadNumber()

	if index < 1 || index > len(p.GetHands()) {
		fmt.Println("請輸入有效的卡牌編號")
		return p.Show()
	}

	card := p.GetHands()[index-1]
	p.ReplaceHands(append(p.GetHands()[:index-1], p.GetHands()[index:]...))

	return card
}

func (p *HumanShowdownPlayer) DecideToExchange() bool {
	fmt.Println("是否要交換卡片？(y/n)")
	result, _ := p.Cli.ReadYesNo()
	return result
}

func (p *HumanShowdownPlayer) UseExchange(players []ShowdownPlayer) (ShowdownPlayer, error) {
	fmt.Printf("輸入想要交換的對象編號(%d ~ %d): \n", 1, len(players))
	for i := 0; i < len(players); i++ {
		fmt.Printf("%d: %s \n", i+1, players[i].GetName())
	}

	index, err := p.Cli.ReadNumber()
	if err != nil {
		return nil, err
	}

	if index < 1 || index > len(players) {
		fmt.Println("請輸入有效的玩家編號")
		return p.UseExchange(players)
	}

	return players[index-1], nil
}

type AIShowdownPlayer struct {
	BaseShowdownPlayer
}

func (p *AIShowdownPlayer) NameSelf() {
	p.setName(fmt.Sprintf("AI Player %d", rand.Intn(1000)))
}

func (p *AIShowdownPlayer) Show() card.CardStrategy {
	if len(p.GetHands()) == 0 {
		fmt.Printf("%s 沒有卡牌可以出\n", p.GetName())
		return nil
	}

	index := rand.Intn(len(p.GetHands()))
	card := p.GetHands()[index]
	p.ReplaceHands(append(p.GetHands()[:index], p.GetHands()[index+1:]...))
	return card
}

func (p *AIShowdownPlayer) DecideToExchange() bool {
	return rand.Intn(2) == 0
}

func (p *AIShowdownPlayer) UseExchange(players []ShowdownPlayer) (ShowdownPlayer, error) {
	return players[rand.Intn(3)], nil
}

var _ ShowdownPlayer = &HumanShowdownPlayer{}
var _ ShowdownPlayer = &AIShowdownPlayer{}
