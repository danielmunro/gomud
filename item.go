package gomud

import "github.com/jinzhu/gorm"

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

type Item struct {
	gorm.Model
	name        string
	description string
	identifiers []string
	attributes  *Attributes
	position    position
}

func NewItem(name string, description string, identifiers []string) *Item {
	return &Item{
		name:        name,
		description: description,
		identifiers: identifiers,
	}
}

func (i *Item) String() string {
	return i.name
}
