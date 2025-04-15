package main

import "fmt"

type HumanPlayer struct {
	BasePlayer
	cli CLI
}

func (p *HumanPlayer) NameSelf() {
	fmt.Println("請輸入有效的名子:")
	name, _ := p.cli.ReadLine()
	p.setName(name)
}

func (p *HumanPlayer) Show() *Card {
	if len(p.GetHands()) == 0 {
		fmt.Printf("%s 沒有卡牌可以出\n", p.GetName())
		return nil
	}

	fmt.Println("輪到你出牌")
	for i, card := range p.GetHands() {
		fmt.Printf("%d: %s \n", i+1, card.String())
	}
	fmt.Println("請輸入要出牌的卡牌編號:")

	index, _ := p.cli.ReadNumber()

	if index < 1 || index > len(p.GetHands()) {
		fmt.Println("請輸入有效的卡牌編號")
		return p.Show()
	}

	card := p.GetHands()[index-1]
	p.setCards(append(p.GetHands()[:index-1], p.GetHands()[index:]...))

	return &card
}

var _ Player = &HumanPlayer{}
