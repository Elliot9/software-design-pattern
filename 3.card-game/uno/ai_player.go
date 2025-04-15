package main

import (
	"fmt"
	"math/rand"
)

type AIPlayer struct {
	BasePlayer
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
	p.setCards(append(p.GetHands()[:index], p.GetHands()[index+1:]...))
	return &card
}

var _ Player = &AIPlayer{}
