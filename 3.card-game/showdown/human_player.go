package main

import "fmt"

type HumanPlayer struct {
	BasePlayer
	cli CLI
}

func (p *HumanPlayer) DecideToExchange() bool {
	fmt.Println("是否要交換卡片？(y/n)")
	result, _ := p.cli.ReadYesNo()
	return result
}

func (p *HumanPlayer) UseExchange(players []Player) (Player, error) {
	fmt.Printf("輸入想要交換的對象編號(%d ~ %d): \n", 1, len(players))
	for i := 0; i < len(players); i++ {
		fmt.Printf("%d: %s \n", i+1, players[i].GetName())
	}

	index, err := p.cli.ReadNumber()
	if err != nil {
		return nil, err
	}

	if index < 1 || index > len(players) {
		fmt.Println("請輸入有效的玩家編號")
		return p.UseExchange(players)
	}

	return players[index-1], nil
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
	p.ReplaceHands(append(p.GetHands()[:index-1], p.GetHands()[index:]...))

	return &card
}

var _ Player = &HumanPlayer{}
