package model

import (
	"errors"
	"github.com/danielmunro/gomud/io"
	"strings"

	"github.com/jinzhu/gorm"
)

type Gender string

const (
	AnyGender Gender = "any"
	NoGender Gender = "none"
	AllGenders Gender = "all"
	FemaleGender Gender = "female"
	MaleGender Gender = "male"
)

type Role string

const (
	Scavenger Role = "scavenger"
	Mobile    Role = "mobile"
	Merchant  Role = "merchant"
)

type Disposition string

const (
	DeadDisposition Disposition = "dead"
	IncapacitatedDisposition = "incapacitated"
	StunnedDisposition = "stunned"
	SleepingDisposition = "sleeping"
	SittingDisposition = "sitting"
	FightingDisposition = "fighting"
	StandingDisposition = "standing"
)

type Mob struct {
	gorm.Model
	Name        string
	description string
	Identifiers []string
	attributes  *Attributes
	Disposition Disposition
	level       int
	Hp          int
	Mana        int
	Mv          int
	race        *Race
	job         *Job
	room        *Room
	lastRoom    *Room
	roles       []Role
	Items       []*Item
	Equipped    []*Item
	Gender      Gender
}

func NewMob(n string, d string) *Mob {
	return &Mob{
		Name:        n,
		description: d,
		Identifiers: strings.Split(n, " "),
		attributes:  NewStartingAttrs(),
		Disposition: StandingDisposition,
		level:       1,
		race:        getRace(CritterRace),
		job:         getJob(UninitializedJob),
		Gender:      NoGender,
	}
}

func (m *Mob) FindItem(b *io.Buffer) (*Item, error) {
	for _, i := range m.Items {
		if b.MatchesSubject(i.identifiers) {
			return i, nil
		}
	}
	return nil, errors.New("no item found")
}

func (m *Mob) FindEquipped(b *io.Buffer) (*Item, error) {
	for _, i := range m.Equipped {
		if b.MatchesSubject(i.identifiers) {
			return i, nil
		}
	}
	return nil, errors.New("no equipment found")
}

func (m *Mob) String() string {
	return m.Name
}

func (m *Mob) SetFightDisposition() {
	m.Disposition = FightingDisposition
}

func (m *Mob) SetIncapacitatedDisposition() {
	m.Disposition = IncapacitatedDisposition
}

func (m *Mob) SetDeadDisposition() {
	m.Disposition = DeadDisposition
}

func (m *Mob) SetSittingDisposition() {
	m.Disposition = SittingDisposition
}

func (m *Mob) SetSleepingDisposition() {
	m.Disposition = SleepingDisposition
}

func (m *Mob) SetStandingDisposition() {
	m.Disposition = StandingDisposition
}

func (m *Mob) CanContinueFighting() bool {
	return m.Disposition != DeadDisposition && m.Disposition != IncapacitatedDisposition
}

func (m *Mob) IsDead() bool {
	return m.Disposition == DeadDisposition
}

func (m *Mob) IsFighting() bool {
	return m.Disposition == FightingDisposition
}

func (m *Mob) IsSleeping() bool {
	return m.Disposition == SleepingDisposition
}

func (m *Mob) HasDisposition(disposition Disposition) bool {
	return m.Disposition == disposition
}

func (m *Mob) Attr(a Attribute) int {
	return m.attributes.Value(a) + m.race.Attributes.Value(a) + m.job.Attributes.Value(a)
}

func (m *Mob) GetGenderPronoun() string {
	switch m.Gender {
	case FemaleGender:
		return "her"
	case MaleGender:
		return "his"
	}
	return "their"
}

func (m *Mob) IsMerchant() bool {
	return m.hasRole(Merchant)
}

func (m *Mob) hasRole(role Role) bool {
	for _, r := range m.roles {
		if r == role {
			return true
		}
	}

	return false
}
