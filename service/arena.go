package service

import (
	"awesomeProject58/model"
	"fmt"
)

type Arena struct {
	Player1 model.Player
	Player2 model.Player
}

func (a *Arena) StartMatch() {
	// Determine starting player
	currentPlayer := a.getStartingPlayer()

	// Match loop
	for {
		damage := currentPlayer.Attacking() // Assuming Attack() is a method on Player struct
		opponent := a.getOpponent(currentPlayer)
		opponent.Defend(damage)

		// Print round results (using indices if Name field is not available)
		fmt.Printf("Player %d attacks Player %d for %d damage.\n", 1, 2, damage)
		fmt.Printf("Player %d's health is now %d.\n", 2, opponent.Health)

		if !opponent.IsAlive() {
			break
		}

		// Switch players
		currentPlayer = opponent
	}
}

func (a *Arena) PrintMatchResults() {
	if a.Player1.IsAlive() {
		fmt.Printf("Player 1 wins! Player 2 has been defeated.\n")
	} else {
		fmt.Printf("Player 2 wins! Player 1 has been defeated.\n")
	}
	fmt.Printf("Player 1's final health: %d\n", a.Player1.Health)
	fmt.Printf("Player 2's final health: %d\n", a.Player2.Health)
}

func (a *Arena) getStartingPlayer() *model.Player {
	if a.Player2.Health < a.Player1.Health {
		return &a.Player2
	}
	return &a.Player1
}

func (a *Arena) getOpponent(player *model.Player) *model.Player {
	if player == &a.Player1 {
		return &a.Player2
	} else {
		return &a.Player1
	}
}
