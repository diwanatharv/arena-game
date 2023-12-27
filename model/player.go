package model

var attackDie = Die{Sides: 6}
var defendDie = Die{Sides: 6}

type Player struct {
	Health, Strength, Attack int
}

// Attacking simulates an attacking action by a player and returns the damage caused
func (p *Player) Attacking() int {
	roll := attackDie.Roll()
	damage := p.Attack * roll
	return damage
}

// Defend simulates a defending action by a player, updating their health based on the damage received.
func (p *Player) Defend(damage int) {
	roll := defendDie.Roll()
	defendedDamage := p.Strength * roll
	p.Health -= damage - defendedDamage
}

// IsAlive checks if the player is still alive based on their health.
func (p *Player) IsAlive() bool {
	return p.Health > 0
}
