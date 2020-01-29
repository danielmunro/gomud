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
	f.Attacker.attack(f.Defender)
	f.Defender.attack(f.Attacker)
	if f.Attacker.hp < 0 || f.Defender.hp < 0 {
		f.End()
	}
}

func (f *Fight) IsEnded() bool {
	return f.Status == EndedFightStatus
}
