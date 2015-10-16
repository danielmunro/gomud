package gomud

import (
	"github.com/jinzhu/gorm"
)

// Position is a string representing where an item can/is be applied
// to a Mob. i.e. "head" or "right hand"
type Position string

// The default five Positions for items.
const (
	Head      Position = "head"
	Torso     Position = "torso"
	Legs      Position = "legs"
	RightHand Position = "right hand"
	LeftHand  Position = "left hand"
)

// Item is an interface for objects within the game world.
// String() - converts the Item to a string.
type Item struct {
	gorm.Model
	ShortName, LongName string
	Attributes          *Attributes
	Position            Position
	Items               []Item
}
