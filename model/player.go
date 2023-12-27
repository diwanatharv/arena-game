package model

type Player struct {
	Health, Strength, Attack int
}

func (p *Player) Attacking() int {
	attackDie := Die{6}
	roll := attackDie.Roll()
	damage := p.Attack * roll
	return damage
}

func (p *Player) Defend(damage int) {
	defendDie := Die{6}
	roll := defendDie.Roll()
	defendedDamage := p.Strength * roll
	p.Health -= damage - defendedDamage
}

func (p *Player) IsAlive() bool {
	return p.Health > 0
}
