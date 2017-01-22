package gomud

import (
	"fmt"

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

type mob struct {
	gorm.Model
	name        string
	description string
	room        *room
}

type exit struct {
	gorm.Model
	room      *room
	direction direction
}

type room struct {
	gorm.Model
	name        string
	description string
	exits       []*exit
	mobs        []*mob `gorm:"ForeignKey:mob"`
}

func newRoom(n string, d string) *room {
	return &room{
		name:        n,
		description: d,
	}
}

// ExitsString lists exits with first character of the direction
func (r *room) exitsString() string {
	var exits string

	for _, e := range r.exits {
		exits = fmt.Sprintf("%s%s", exits, string(e.direction[0]))
	}

	return fmt.Sprintf("[%s]", exits)
}

func (r *room) mobsString() string {
	var mobs string

	for _, m := range r.mobs {
		mobs = fmt.Sprintf("%s is here.\n%s", string(m.name), mobs)
	}

	return mobs
}
