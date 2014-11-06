package gomud

type Position string

const (
	Head      Position = "head"
	Torso     Position = "torso"
	Legs      Position = "legs"
	RightHand Position = "right hand"
	LeftHand  Position = "left hand"
)

type Item interface {
	String() string
}

type Equipment struct {
	ShortName, LongName string
	Attributes          *Attributes
	Position            Position
}

func (e Equipment) String() string {
	return e.ShortName
}

type Container struct {
	Equipment
	Items []Item
}

type Equipped struct {
	Head, Torso, Legs, RightHand, LeftHand *Equipment
}

func (e Equipped) getAll() map[Position]*Equipment {
	return map[Position]*Equipment{
		Head:      e.Head,
		Torso:     e.Torso,
		Legs:      e.Legs,
		RightHand: e.RightHand,
		LeftHand:  e.LeftHand,
	}
}
