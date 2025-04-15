package main

import (
	"fmt"
	"math/rand"
)

type AIPlayer struct {
	BasePlayer
}

func (p *AIPlayer) DecideToExchange() bool {
	return rand.Intn(2) == 0
}

func (p *AIPlayer) UseExchange(players []Player) (Player, error) {
	if p.UsedPermission() {
		return nil, fmt.Errorf("已經使用過交換卡牌")
	}

	return players[rand.Intn(3)], nil
}

func (p *AIPlayer) NameSelf() {
	p.setName(fmt.Sprintf("AI Player %d", rand.Intn(1000)))
}

func (p *AIPlayer) Show() *Card {
	if len(p.GetHands()) == 0 {
		fmt.Printf("%s 沒有卡牌可以出\n", p.GetName())
		return nil
	}

	// 隨機出牌
	index := rand.Intn(len(p.GetHands()))
	card := p.GetHands()[index]
	p.ReplaceHands(append(p.GetHands()[:index], p.GetHands()[index+1:]...))
	return &card
}

var _ Player = &AIPlayer{}
