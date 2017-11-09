package main

import (
	. "codewars"
	"codewars/runner"
)

func main() {
	runner.Start(&MyStrategy{})
}

type MyStrategy struct {
}

func (m *MyStrategy) Move(*Player, *World, *Game, *Move) {

}
