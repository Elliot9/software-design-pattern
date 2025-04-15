package main

import "fmt"

// 表示一張撲克牌，供階級花色比較
type Card struct {
	rank Rank
	suit Suit
}

func (c *Card) getRank() Rank {
	return c.rank
}

func (c *Card) getSuit() Suit {
	return c.suit
}

func (c *Card) String() string {
	return fmt.Sprintf("%s %s", c.suit.string(), c.rank.string())
}

type Rank int

func (r Rank) string() string {
	return []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}[r]
}

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

func (s Suit) string() string {
	return []string{"♣", "♦", "♥", "♠"}[s]
}
