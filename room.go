package gomud

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type direction string

const (
	dNorth direction = "north"
	dSouth direction = "south"
	dEast  direction = "east"
	dWest  direction = "west"
	dUp    direction = "up"
	dDown  direction = "down"
)

func reverseDirection(d direction) (direction, error) {
	switch d {
	case dNorth:
		return dSouth, nil
	case dSouth:
		return dNorth, nil
	case dEast:
		return dWest, nil
	case dWest:
		return dEast, nil
	case dUp:
		return dDown, nil
	case dDown:
		return dUp, nil
	default:
		return d, fmt.Errorf("unexpected direction %s", string(d))
	}
}

type exit struct {
	gorm.Model
	room      *Room
	direction direction
}

func newExit(r *Room, d direction) *exit {
	return &exit{
		room:      r,
		direction: d,
	}
}

type Room struct {
	gorm.Model
	name        string
	description string
	exits       []*exit
	mobs        []*Mob `gorm:"ForeignKey:Mob"`
	items       []*item
}

func newRoom(n string, d string) *Room {
	return &Room{
		name:        n,
		description: d,
		items:       []*item{},
	}
}

func (r *Room) String() string {
	return r.name
}
