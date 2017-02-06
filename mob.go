package gomud

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type role string

const (
	scavenger role = "scavenger"
	mobile    role = "mobile"
)

type mob struct {
	gorm.Model
	name        string
	description string
	identifiers []string
	attributes  *attributes
	level       int
	hp          int
	mana        int
	mv          int
	race        race
	job         job
	room        *room
	lastRoom    *room
	roles       []role
	client      *client
	items       []*item
	equipped    []*item
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
	message := fmt.Sprintf("%s leaves heading %s.\n", m.String(), e.direction)
	for i, rm := range m.room.mobs {
		if rm == m {
			m.room.mobs = append(m.room.mobs[0:i], m.room.mobs[i+1:]...)
		} else {
			rm.notify(message)
		}
	}
	m.room = e.room
	m.room.mobs = append(m.room.mobs, m)
	message = fmt.Sprintf("%s arrives.\n", m.String())
	for _, rm := range m.room.mobs {
		if rm != m {
			rm.notify(message)
		}
	}
	log.Println(fmt.Sprintf("%s moves to %s", m.String(), m.room.String()))
}

func (m *mob) roam() {
	c := len(m.room.exits)
	if c == 0 {
		return
	} else if c == 1 {
		m.move(m.room.exits[0])
	} else {
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
		newAction(m, fmt.Sprintf("get %s", m.room.items[0].identifiers[0]))
	}
}
