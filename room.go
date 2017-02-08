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
