package tests

import (
	"awesomeProject58/model"
	"testing"
)

func TestDieRoll(t *testing.T) {
	die := model.Die{Sides: 6}
	for i := 0; i < 100; i++ {
		roll := die.Roll()
		if roll < 1 || roll > die.Sides {
			t.Errorf("Die roll outside expected range: %d", roll)
		}
	}
}
