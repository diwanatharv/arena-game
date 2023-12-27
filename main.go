package main

import (
	"awesomeProject58/model"
	"awesomeProject58/service"
)

func main() {
	player1 := model.Player{Health: 50, Strength: 5, Attack: 10}
	player2 := model.Player{Health: 100, Strength: 10, Attack: 5}
	arena := service.Arena{Player1: player1, Player2: player2}
	arena.StartMatch()
	arena.PrintMatchResults()
}
