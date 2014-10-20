package gomud

type Item struct {
	ShortName, LongName string
	Attributes          *Attributes
}

type Equipment string

const (
	Head      Equipment = "head"
	Torso     Equipment = "torso"
	Legs      Equipment = "legs"
	RightHand Equipment = "right hand"
	LeftHand  Equipment = "left hand"
)

type Equipped struct {
	Head, Torso, Legs, RightHand, LeftHand *Item
}

func (e Equipped) getAll() map[Equipment]*Item {
	return map[Equipment]*Item{
		Head:      e.Head,
		Torso:     e.Torso,
		Legs:      e.Legs,
		RightHand: e.RightHand,
		LeftHand:  e.LeftHand,
	}
}
