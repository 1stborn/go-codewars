package main

import (
	"codewars"
	"codewars/model"
	"os"
)

type MyStrategy struct {
}

func main() {
	codewars.Start(new(MyStrategy), os.Args[1:]...)
}

func (m *MyStrategy) NewGame() {
}

func (m *MyStrategy) Move(move *runner.Move) {

}
