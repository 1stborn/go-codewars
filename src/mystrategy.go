package main

import . "codewars"

type MyStrategy struct {
}

func New() *MyStrategy {
	return new(MyStrategy)
}

func (m *MyStrategy) Move(player *Player, world *World, game *Game, move *Move) {
}
