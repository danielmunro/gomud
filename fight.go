package gomud

import "fmt"

type fight struct {
	m1 *mob
	m2 *mob
}

func newFight(m1 *mob, m2 *mob) *fight {
	m1.disposition = fightingDisposition
	m2.disposition = fightingDisposition
	m1.notify(fmt.Sprintf("You scream and attack %s!", m2.String()))

	f := &fight{
		m1: m1,
		m2: m2,
	}

	f.turn(m1)

	m1.fight = f
	m2.fight = f

	return f
}

func (f *fight) turn(m *mob) {
	if m == f.m1 {
		m.attack(f.m2)
	} else if m == f.m2 {
		m.attack(f.m1)
	}
}
