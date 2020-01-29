package gomud

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

type disposition int

const (
	deadDisposition disposition = iota
	incapacitatedDisposition
	stunnedDisposition
	sleepingDisposition
	sittingDisposition
	fightingDisposition
	standingDisposition
)

type Mob struct {
	gorm.Model
	name        string
	description string
	identifiers []string
	attributes  *attributes
	disposition disposition
	level       int
	hp          int
	mana        int
	mv          int
	race        *race
	job         job
	room        *Room
	lastRoom    *Room
	roles       []role
	items       []*item
	equipped    []*item
}

func newMob(n string, d string) *Mob {
	return &Mob{
		name:        n,
		description: d,
		identifiers: strings.Split(n, " "),
		attributes:  &attributes{},
		disposition: standingDisposition,
		level:       1,
		race:        getRace(""),
		job:         uninitiated,
	}
}

func (m *Mob) FindItem(b *io.Buffer) (*item, error) {
	for _, i := range m.items {
		if b.MatchesSubject(i.identifiers) {
			return i, nil
		}
	}
	return nil, errors.New("no item found")
}

func (m *Mob) FindEquipped(b *io.Buffer) (*item, error) {
	for _, i := range m.equipped {
		if b.MatchesSubject(i.identifiers) {
			return i, nil
		}
	}
	return nil, errors.New("no equipment found")
}

func (m *Mob) String() string {
	return m.name
}

func (m *Mob) hasRole(r role) bool {
	for _, mr := range m.roles {
		if mr == r {
			return true
		}
	}

	return false
}

func (m *Mob) attr(a attribute) int {
	return m.attributes.a(a) + m.race.attrs.a(a) + jobAttributes(m.job).a(a)
}

func (m *Mob) attack(target *Mob) {
	if target.disposition > deadDisposition {
		target.hp -= dice().Intn(m.attr(aDam)) + m.attr(aHit)
	}
}
