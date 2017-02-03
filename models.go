package gomud

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type position string

const (
	light     position = "light"
	finger1   position = "finger"
	finger2   position = "finger"
	neck1     position = "neck"
	neck2     position = "neck"
	torso     position = "torso"
	head      position = "head"
	legs      position = "legs"
	feet      position = "feet"
	hands     position = "hands"
	arms      position = "arms"
	shield    position = "shield"
	body      position = "body"
	waist     position = "waist"
	wrist1    position = "wrist"
	wrist2    position = "wrist"
	wield     position = "wield"
	held      position = "held"
	floating  position = "floating"
	secondary position = "secondary"
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

func reverseDirection(d direction) (direction, error) {
	switch d {
	case north:
		return south, nil
	case south:
		return north, nil
	case east:
		return west, nil
	case west:
		return east, nil
	case up:
		return down, nil
	case down:
		return up, nil
	default:
		return d, errors.New(fmt.Sprintf("unexpected direction %s", string(d)))
	}
}

type role string

const (
	scavenger role = "scavenger"
	mobile    role = "mobile"
)

type event struct {
	sender  *mob
	message string
}

func newEvent(s *mob, m string) *event {
	return &event{
		sender:  s,
		message: m,
	}
}

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
		get(&input{
			args: []string{"get", m.room.items[0].identifiers[0]},
			mob:  m,
		})
	}
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
		items:       []*item{},
	}
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
	acBash   int
	acSlash  int
	acPierce int
	acMagic  int
}

type item struct {
	gorm.Model
	name        string
	description string
	identifiers []string
	attributes  *attributes
	position    position
}

func newItem(name string, description string, identifiers []string) *item {
	return &item{
		name:        name,
		description: description,
		identifiers: identifiers,
	}
}

func (i *item) String() string {
	return i.name
}
