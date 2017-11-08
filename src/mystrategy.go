package main

import (
	"codewars"
	"os"
	"codewars/model"
)

type MyStrategy struct {
	*model.World
}

func main() {
	codewars.Start(new(MyStrategy), os.Args[1:]...)
}

func (m *MyStrategy) NewGame(game *model.Game)  {

}

func (m *MyStrategy) Move(move *model.Move)  {

}

func (m *MyStrategy) GetWorld() *model.World {
	return m.World
}
