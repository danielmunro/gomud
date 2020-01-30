package gomud

import (
	"github.com/danielmunro/gomud/model"
	"github.com/danielmunro/gomud/util"
)

type FightStatus string

const (
	StartedFightStatus FightStatus = "started"
	EndedFightStatus   FightStatus = "ended"
)

type Fight struct {
	Attacker *model.Mob
	Defender *model.Mob
	Status   FightStatus
}

func NewFight(attacker *model.Mob, defender *model.Mob) *Fight {
	attacker.SetFightDisposition()
	defender.SetFightDisposition()
	return &Fight{
		Attacker: attacker,
		Defender: defender,
		Status:   StartedFightStatus,
	}
}

func (f *Fight) IncludesMob(mob *model.Mob) bool {
	return f.Attacker == mob || f.Defender == mob
}

func (f *Fight) End() {
	f.Status = EndedFightStatus
}

func (f *Fight) Proceed() {
	attack(f.Attacker, f.Defender)
	attack(f.Defender, f.Attacker)
	if !f.Attacker.CanContinueFighting() || !f.Defender.CanContinueFighting() {
		f.End()
	}
}

func (f *Fight) IsEnded() bool {
	return f.Status == EndedFightStatus
}

func attack(attacker *model.Mob, defender *model.Mob) {
	if attacker.IsFighting() && !defender.IsDead() {
		defender.Hp -= util.Dice().Intn(attacker.Attr(model.DamAttr)) + attacker.Attr(model.HitAttr)
		if defender.Hp < 20 {
			defender.SetDeadDisposition()
		} else if defender.Hp < 0 {
			defender.SetIncapacitatedDisposition()
		}
	}
}
