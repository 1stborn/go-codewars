package main

import . "codewars"

type MyStrategy struct {
}

func New() *MyStrategy {
	return new(MyStrategy)
}

func (m *MyStrategy) Move(player *Player, world *World, game *Game, move *Move) {
	if world.TickIndex == 0 {
		move.Action = Action_ClearAndSelect
		move.Right = world.Width
		move.Bottom = world.Height
		return
	}

	if world.TickIndex == 1 {
		move.Action = Action_Move
		move.X = world.Width / 2.0
		move.Y = world.Height / 2.0
	}
}
