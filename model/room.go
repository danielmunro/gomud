package model

import (
	"errors"
	"fmt"
	"github.com/danielmunro/gomud/io"

	"github.com/jinzhu/gorm"
)

type Direction string

const (
	NorthDirection Direction = "north"
	SouthDirection Direction = "south"
	EastDirection  Direction = "east"
	WestDirection  Direction = "west"
	UpDirection    Direction = "up"
	DownDirection  Direction = "down"
)

func reverseDirection(d Direction) (Direction, error) {
	switch d {
	case NorthDirection:
		return SouthDirection, nil
	case SouthDirection:
		return NorthDirection, nil
	case EastDirection:
		return WestDirection, nil
	case WestDirection:
		return EastDirection, nil
	case UpDirection:
		return DownDirection, nil
	case DownDirection:
		return UpDirection, nil
	default:
		return d, fmt.Errorf("unexpected direction %s", string(d))
	}
}

type Exit struct {
	gorm.Model
	Room      *Room
	Direction Direction
}

func NewExit(r *Room, d Direction) *Exit {
	return &Exit{
		Room:      r,
		Direction: d,
	}
}

type Room struct {
	gorm.Model
	Name        string
	Description string
	Exits       []*Exit
	Items       []*Item
}

func NewRoom(n string, d string) *Room {
	return &Room{
		Name:        n,
		Description: d,
		Items:       []*Item{},
	}
}

func (r *Room) GetExitFromDirection(direction string) (*Exit, error) {
	for _, e := range r.Exits {
		if string(e.Direction) == direction {
			return e, nil
		}
	}
	return nil, errors.New("that direction does not exist")
}

func (r *Room) FindItem(b *io.Buffer) (*Item, error) {
	for _, i := range r.Items {
		if b.MatchesSubject(i.identifiers) {
			return i, nil
		}
	}
	return nil, errors.New("no item found")
}

func (r *Room) String() string {
	return r.Name
}
