package gomud

/*
	Position is a string representing where an item can/is be applied
	to a Mob. i.e. "head" or "right hand"
*/
type Position string

/*
	The default five Positions for items.
*/
const (
	Head      Position = "head"
	Torso     Position = "torso"
	Legs      Position = "legs"
	RightHand Position = "right hand"
	LeftHand  Position = "left hand"
)

/*
	Item is an interface for objects within the game world.
	String() - converts the Item to a string.
*/
type Item interface {
	String() string
}

/*
	Equipment represents an Item that can be "equipped" to
	a Mob's person.
	ShortName - a short string describing the item.
	LongName - a long string describing the item.
	Attributes - the properties of the item, represented as an Attributes struct.
	Position - the place that an item can be equipped to a Mob.
*/
type Equipment struct {
	ShortName, LongName string
	Attributes          *Attributes
	Position            Position
}

/*
	String returns the Equipment represented as a string.
*/
func (e Equipment) String() string {
	return e.ShortName
}

/*
	Container represents a collection of items.
	Equipment - an equipment representing the Container itself in the game context.
	Items - an array of Items within the container.
*/
type Container struct {
	Equipment
	Items []Item
}

/*
	Equipped contains an equipment for each default Position.
*/
type Equipped struct {
	Head, Torso, Legs, RightHand, LeftHand *Equipment
}

/*
	getAll returns a map of each position in an Equipped to
	the Equipment at that position.
*/
func (e Equipped) getAll() map[Position]*Equipment {
	return map[Position]*Equipment{
		Head:      e.Head,
		Torso:     e.Torso,
		Legs:      e.Legs,
		RightHand: e.RightHand,
		LeftHand:  e.LeftHand,
	}
}
