package gomud

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
)

/*
	Direction is a string indicating a direction of movement.
	i.e. "north" or "up"
*/
type Direction string

/*
	Six common pre-defined Directions.
*/
const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
	Up    Direction = "up"
	Down  Direction = "down"
)

/*
	Room represents a single location within the MUD. Its attributes represent
	how it is connected to other rooms, who and what is present, and what it looks
	like.
	Title - the name of the location as a string.
	Description - what the room looks like as a string.
	Area - the larger region of the world in which the room is located.
	Directions - a map of Directions and Room identification numbers.
	Rooms - a map of Directions to associated Rooms.
	Mobs - an array of the Mobs currently within the Room.
	MovementCost - the quantity of movement that it takes to enter the room.
	Items - an array of Items currently at this location.
*/
type Room struct {
	Title, Description, Area string
	Directions               map[Direction]int
	Rooms                    map[Direction]*Room
	Mobs                     []*Mob
	MovementCost             float64
	Items                    []Item
}

/*
	init populates the rooms from a specified file.
*/
func init() {
	dir, _ := filepath.Abs(filepath.Dir("areas/"))
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

/*
	rooms maps Room numbers to individual Rooms.
*/
var rooms map[int]*Room

/*
	AddMob inserts a given Mob m into the room.
*/
func (r *Room) AddMob(m *Mob) {
	r.Mobs = append(r.Mobs, m)
}

/*
	RemoveMob removes Mob m from the Room.
*/
func (r *Room) RemoveMob(m *Mob) {
	for p, v := range r.Mobs {
		if v == m {
			r.Mobs = append(r.Mobs[0:p], r.Mobs[p+1:]...)
			return
		}
	}
}

/*
	AllDirections returns an array of all valid exit Directions from
	the Room.
*/
func (r *Room) AllDirections() (dirs []Direction) {
	for d, _ := range r.Rooms {
		dirs = append(dirs, d)
	}
	return dirs
}

/*
	Finds a mob by the name arg in the room.
*/
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

/*
	Announce notifies a Mob within a room of a given message.
*/
func (r *Room) Announce(m *Mob, message string) {
	for _, mob := range r.Mobs {
		if m != mob {
			mob.Notify(message)
		}
	}
}

/*
	OppositeDirection returns the opposite of a given Direction.
	i.e. OppositeDirection("north") would return "south"
*/
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
