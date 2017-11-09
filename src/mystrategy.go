package main

import (
	"codewars"
	"codewars/model"
	"os"
)

type MyStrategy struct {
	*model.Game
	*model.World
}

func main() {
	codewars.Start(new(MyStrategy), os.Args[1:]...)
}

func (m *MyStrategy) NewGame(game *model.Game) {
	m.Game = game
	m.World = &model.World{
		Players:    make(map[int64]*model.Player),
		Vehicles:   make(map[int64]*model.Vehicle),
		Facilities: make(map[int64]*model.Facility),
	}
}

func (m *MyStrategy) Move(move *model.Move) {

}

func (m *MyStrategy) GetWorld() *model.World {
	return m.World
}
