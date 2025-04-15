package main

import "github/elliot9/card_game_template/game"

var g game.Game

func main() {
	// g = game.NewUnoGame()
	g = game.NewShowdownGame()
	g.Start()
}
