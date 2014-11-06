package gomud

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
	Up    Direction = "up"
	Down  Direction = "down"
)

type Room struct {
	Title, Description, Area string
	Directions               map[Direction]int
	Rooms                    map[Direction]*Room
	Mobs                     []*Mob
	MovementCost             float64
	Items                    []Item
}

func init() {
	dir, _ := filepath.Abs(filepath.Dir("gomud/areas/"))
	data, _ := ioutil.ReadFile(dir + "/midgaard.yaml")
	yaml.Unmarshal(data, &rooms)
	for _, r := range rooms {
		r.Rooms = make(map[Direction]*Room, len(r.Directions))
		for d, roomId := range r.Directions {
			r.Rooms[d] = rooms[roomId]
		}
		for _, m := range r.Mobs {
			m.CurrentAttr = &Attributes{}
			*m.CurrentAttr = *m.Attributes
			m.Room = r
			mobs = append(mobs, m)
		}
	}
}

var rooms map[int]*Room

func (r *Room) AddMob(m *Mob) {
	r.Mobs = append(r.Mobs, m)
}

func (r *Room) RemoveMob(m *Mob) {
	for p, v := range r.Mobs {
		if v == m {
			r.Mobs = append(r.Mobs[0:p], r.Mobs[p+1:]...)
			return
		}
	}
}

func (r *Room) AllDirections() (dirs []Direction) {
	for d, _ := range r.Rooms {
		dirs = append(dirs, d)
	}
	return dirs
}

func (r *Room) FindMob(arg string) *Mob {

	for _, m := range r.Mobs {
		names := strings.Split(m.ShortName, " ")
		for _, n := range names {
			if strings.Index(strings.ToLower(n), arg) == 0 {
				return m
			}
		}
	}

	return nil
}

func (r *Room) Announce(m *Mob, message string) {
	for _, mob := range r.Mobs {
		if m != mob {
			mob.Notify(message)
		}
	}
}

func OppositeDirection(d Direction) (Direction, error) {
	switch d {
	case North:
		return South, nil
	case South:
		return North, nil
	case East:
		return West, nil
	case West:
		return East, nil
	case Up:
		return Down, nil
	case Down:
		return Up, nil
	default:
		return "", errors.New("Not a valid direction")
	}
}
