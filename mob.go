package gomud

import (
	"fmt"
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

type mob struct {
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
	room        *room
	lastRoom    *room
	roles       []role
	client      *client
	items       []*item
	equipped    []*item
	fight       *fight
}

func newMob(n string, d string) *mob {
	return &mob{
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

func (m *mob) notify(message string) {
	if m.client != nil {
		m.client.writePrompt(message)
	}
}

func (m *mob) String() string {
	return m.name
}

func (m *mob) hasRole(r role) bool {
	for _, mr := range m.roles {
		if mr == r {
			return true
		}
	}

	return false
}

func (m *mob) move(e *exit) {
	m.lastRoom = m.room
	for i, rm := range m.room.mobs {
		if rm == m {
			m.room.mobs = append(m.room.mobs[0:i], m.room.mobs[i+1:]...)
		} else {
			rm.notify(fmt.Sprintf("%s leaves heading %s.\n", m.String(), e.direction))
		}
	}
	m.room = e.room
	m.room.mobs = append(m.room.mobs, m)
	for _, rm := range m.room.mobs {
		if rm != m {
			rm.notify(fmt.Sprintf("%s arrives.\n", m.String()))
		}
	}
}

func (m *mob) roam() {
	switch c := len(m.room.exits); c {
	case 0:
		return
	case 1:
		m.move(m.room.exits[0])
	default:
		for {
			e := m.room.exits[dice().Intn(c)]
			if e.room != m.lastRoom {
				m.move(e)
				break
			}
		}
	}
}

func (m *mob) scavenge() {
	if len(m.room.items) > 0 {
		//newActionWithMob(m, fmt.Sprintf("get %s", m.room.items[0].identifiers[0]))
	}
}

func (m *mob) attr(a attribute) int {
	return m.attributes.a(a) + m.race.attrs.a(a) + jobAttributes(m.job).a(a)
}

func (m *mob) attack(target *mob) {
	if target.disposition > deadDisposition {
		target.hp -= dice().Intn(m.attr(aDam)) + m.attr(aHit)
	}
}
