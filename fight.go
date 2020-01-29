package gomud

type FightStatus string

const (
	StartedFightStatus FightStatus = "started"
	EndedFightStatus   FightStatus = "ended"
)

type Fight struct {
	Attacker *Mob
	Defender *Mob
	Status   FightStatus
}

func NewFight(attacker *Mob, defender *Mob) *Fight {
	attacker.disposition = fightingDisposition
	defender.disposition = fightingDisposition
	return &Fight{
		Attacker: attacker,
		Defender: defender,
		Status:   StartedFightStatus,
	}
}

func (f *Fight) IncludesMob(mob *Mob) bool {
	return f.Attacker == mob || f.Defender == mob
}

func (f *Fight) End() {
	f.Status = EndedFightStatus
}

func (f *Fight) Proceed() {
	f.turn(f.Attacker)
	f.turn(f.Defender)
	if f.Attacker.hp < 0 || f.Defender.hp < 0 {
		f.End()
	}
}

func (f *Fight) IsEnded() bool {
	return f.Status == EndedFightStatus
}

func (f *Fight) turn(m *Mob) {
	if m == f.Attacker {
		m.attack(f.Defender)
	} else if m == f.Defender {
		m.attack(f.Attacker)
	}
}
