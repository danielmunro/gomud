package gomud

type RaceType string

const (
	Giant RaceType = "giant"
	Human RaceType = "human"
	Elf   RaceType = "elf"
	Dwarf RaceType = "dwarf"
	Gnome RaceType = "gnome"
	Nymph RaceType = "nymph"
)

type Race struct {
	Name, Description string
	Attributes        *Attributes
	Skills            []*Skill
	Playable          bool
}

var races map[RaceType]*Race

func init() {
	races = map[RaceType]*Race{
		Giant: &Race{
			Name:        string(Giant),
			Description: "A giant",
			Playable:    false,
		},
		Human: &Race{
			Name:        string(Human),
			Description: "A human",
			Playable:    true,
		},
		Elf: &Race{
			Name:        string(Elf),
			Description: "An elf",
			Playable:    true,
		},
		Dwarf: &Race{
			Name:        string(Dwarf),
			Description: "A dwarf",
			Playable:    true,
		},
		Gnome: &Race{
			Name:        string(Gnome),
			Description: "A gnome",
			Playable:    true,
		},
		Nymph: &Race{
			Name:        string(Nymph),
			Description: "A nymph",
			Playable:    true,
		},
	}
}
