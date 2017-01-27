package gomud

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type direction string

const (
	north direction = "north"
	south direction = "south"
	east  direction = "east"
	west  direction = "west"
	up    direction = "up"
	down  direction = "down"
)

type role string

const (
	scavenger role = "scavenger"
	mobile    role = "mobile"
)

type notification string

const (
	movement notification = "movement"
)

type event struct {
	sender  *mob
	context notification
	message string
}

func newEvent(s *mob, c notification, m string) *event {
	return &event{
		sender:  s,
		context: c,
		message: m,
	}
}

type mob struct {
	gorm.Model
	name        string
	description string
	attributes  *attributes
	level       int
	hp          int
	mana        int
	mv          int
	room        *room
	lastRoom    *room
	roles       []role
	client      *client
	items       []*item
}

func (m *mob) notify(e *event) {
	if m.client != nil {
		m.client.write(e.message)
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
	v := newEvent(m, movement, fmt.Sprintf("%s leaves.\n", m.String()))
	for i, rm := range m.room.mobs {
		if rm == m {
			m.room.mobs = append(m.room.mobs[0:i], m.room.mobs[i+1:]...)
		} else {
			rm.notify(v)
		}
	}
	m.room = e.room
	m.room.mobs = append(m.room.mobs, m)
	v = newEvent(m, movement, fmt.Sprintf("%s arrives.\n", m.String()))
	for _, rm := range m.room.mobs {
		if rm != m {
			rm.notify(v)
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
}

type exit struct {
	gorm.Model
	room      *room
	direction direction
}

func newExit(r *room, d direction) *exit {
	return &exit{
		room:      r,
		direction: d,
	}
}

type room struct {
	gorm.Model
	name        string
	description string
	exits       []*exit
	mobs        []*mob `gorm:"ForeignKey:mob"`
	items       []*item
}

func newRoom(n string, d string) *room {
	return &room{
		name:        n,
		description: d,
	}
}

func (r *room) exitsString() string {
	var exits string

	for _, e := range r.exits {
		exits = fmt.Sprintf("%s%s", exits, string(e.direction[0]))
	}

	return fmt.Sprintf("[%s]", exits)
}

func (r *room) mobsString(mob *mob) string {
	var mobs string

	for _, m := range r.mobs {
		if m != mob {
			mobs = fmt.Sprintf("%s is here.\n%s", string(m.name), mobs)
		}
	}

	return mobs
}

func (r *room) String() string {
	return r.name
}

type attributes struct {
	gorm.Model
	hp       int
	mana     int
	mv       int
	str      int
	int      int
	wis      int
	dex      int
	con      int
	hit      int
	dam      int
	acPierce int
	acBash   int
	acSlash  int
	acMagic  int
}

type item struct {
	gorm.Model
	name        string
	description string
	attributes  *attributes
}

func newItem(name string, description string) *item {
	return &item{
		name:        name,
		description: description,
	}
}
