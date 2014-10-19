package gomud

type Item struct {
    ShortName, LongName string
    Attributes *Attributes
}

type Equipment int

const (
    Head Equipment = iota
    Torso
    Legs
    RightHand
    LeftHand
)

var slots = [...]string {
    "head",
    "torso",
    "legs",
    "right hand",
    "left hand",
}

func (e Equipment) String() string {
    return slots[e]
}

type Equipped struct {
    Head, Torso, Legs, RightHand, LeftHand *Item
}

func (e Equipped) getAll() map[Equipment]*Item {
    return map[Equipment]*Item {
        Head: e.Head,
        Torso: e.Torso,
        Legs: e.Legs,
        RightHand: e.RightHand,
        LeftHand: e.LeftHand,
    }
}
