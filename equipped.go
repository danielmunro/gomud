package gomud

import (
	"github.com/jinzhu/gorm"
)

// Equipped contains an equipment for each default Position.
type Equipped struct {
	gorm.Model
	Head, Torso, Legs, RightHand, LeftHand *Item
}

// getAll returns a map of each position in an Equipped to
// the Equipment at that position.
func (e Equipped) getAll() map[Position]*Item {
	return map[Position]*Item{
		Head:      e.Head,
		Torso:     e.Torso,
		Legs:      e.Legs,
		RightHand: e.RightHand,
		LeftHand:  e.LeftHand,
	}
}
