package gomud

type fight struct {
	m1 *Mob
	m2 *Mob
}

func newFight(m1 *Mob, m2 *Mob) *fight {
	m1.disposition = fightingDisposition
	m2.disposition = fightingDisposition
	//m1.notify(fmt.Sprintf("You scream and attack %s!", m2.String()))

	f := &fight{
		m1: m1,
		m2: m2,
	}

	f.turn(m1)

	m1.fight = f
	m2.fight = f

	return f
}

func (f *fight) turn(m *Mob) {
	if m == f.m1 {
		m.attack(f.m2)
	} else if m == f.m2 {
		m.attack(f.m1)
	}
}
