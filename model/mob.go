package model

import (
	"errors"
	"github.com/danielmunro/gomud/io"
	"strings"

	"github.com/jinzhu/gorm"
)

type role string

const (
	scavenger role = "scavenger"
	mobile    role = "mobile"
)

type Disposition int

const (
	DeadDisposition Disposition = iota
	IncapacitatedDisposition
	StunnedDisposition
	SleepingDisposition
	SittingDisposition
	FightingDisposition
	StandingDisposition
)

type Mob struct {
	gorm.Model
	Name        string
	description string
	Identifiers []string
	attributes  *Attributes
	disposition Disposition
	level       int
	Hp          int
	Mana        int
	Mv          int
	race        *Race
	job         *Job
	room        *Room
	lastRoom    *Room
	roles       []role
	Items       []*Item
	Equipped    []*Item
}

func NewMob(n string, d string) *Mob {
	return &Mob{
		Name:        n,
		description: d,
		Identifiers: strings.Split(n, " "),
		attributes:  NewStartingAttrs(),
		disposition: StandingDisposition,
		level:       1,
		race:        getRace(CritterRace),
		job:         getJob(UninitializedJob),
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
	m.disposition = FightingDisposition
}

func (m *Mob) SetIncapacitatedDisposition() {
	m.disposition = IncapacitatedDisposition
}

func (m *Mob) SetDeadDisposition() {
	m.disposition = DeadDisposition
}

func (m *Mob) CanContinueFighting() bool {
	return m.disposition > IncapacitatedDisposition
}

func (m *Mob) IsDead() bool {
	return m.disposition == DeadDisposition
}

func (m *Mob) IsFighting() bool {
	return m.disposition == FightingDisposition
}

func (m *Mob) HasDisposition(disposition Disposition) bool {
	return m.disposition == disposition
}

func (m *Mob) Attr(a Attribute) int {
	return m.attributes.Value(a) + m.race.Attributes.Value(a) + m.job.Attributes.Value(a)
}
