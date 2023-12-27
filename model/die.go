package model

import (
	"math/rand"
	"time"
)

type Die struct {
	Sides int
}

func (d *Die) Roll() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(d.Sides) + 1
}
