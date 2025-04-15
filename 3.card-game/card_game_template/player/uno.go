package player

import (
	"fmt"
	"github/elliot9/card_game_template/card"
	"math/rand"
)

type UnoPlayer interface {
	Player
}

type BaseUnoPlayer struct {
	BasePlayer
}

type HumanUnoPlayer struct {
	BaseUnoPlayer
	Cli CLI
}

func (p *HumanUnoPlayer) NameSelf() {
	fmt.Println("請輸入有效的名子:")
	name, _ := p.Cli.ReadLine()
	p.setName(name)
}

func (p *HumanUnoPlayer) Show() card.CardStrategy {
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
	p.setHands(append(p.GetHands()[:index-1], p.GetHands()[index:]...))

	return card
}

type AIUnoPlayer struct {
	BaseUnoPlayer
}

func (p *AIUnoPlayer) NameSelf() {
	p.setName(fmt.Sprintf("AI Player %d", rand.Intn(1000)))
}

func (p *AIUnoPlayer) Show() card.CardStrategy {
	if len(p.GetHands()) == 0 {
		fmt.Printf("%s 沒有卡牌可以出\n", p.GetName())
		return nil
	}

	// 隨機出牌
	index := rand.Intn(len(p.GetHands()))
	card := p.GetHands()[index]
	p.setHands(append(p.GetHands()[:index], p.GetHands()[index+1:]...))
	return card
}

var _ UnoPlayer = &HumanUnoPlayer{}
var _ UnoPlayer = &AIUnoPlayer{}
