package model

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type Die struct {
	Sides int
}

// Roll simulates rolling the die and returns a random number between 1 and the number of sides.
func (d *Die) Roll() int {

	return random.Intn(d.Sides) + 1
}
