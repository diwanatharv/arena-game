package tests

import (
	"awesomeProject58/model"
	"testing"
)

func TestPlayerIsAlive(t *testing.T) {
	alivePlayer := model.Player{Health: 10}
	deadPlayer := model.Player{Health: 0}

	if !alivePlayer.IsAlive() {
		t.Errorf("Alive player should be marked alive")
	}

	if deadPlayer.IsAlive() {
		t.Errorf("Dead player should be marked dead")
	}
}
