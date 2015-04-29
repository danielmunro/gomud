package gomud

/*
	RaceType is a string representing the name of a fantasy race.
*/
type RaceType string

/*
	Six default RaceTypes.
*/
const (
	Giant RaceType = "giant"
	Human RaceType = "human"
	Elf   RaceType = "elf"
	Dwarf RaceType = "dwarf"
	Gnome RaceType = "gnome"
	Nymph RaceType = "nymph"
)

/*
	Race defines the attributes of a fantasy race.

	Name - a string name for the race.
	Description - a string describing the race.
	Attributes - an Attributes struct defining the maximum Attributes for the race.
	Skills - an array of Skills for the race.
	Playable - a boolean value indicating whether players can be members of
				the race.
*/
type Race struct {
	Name, Description string
	Attributes        *Attributes
	Skills            []*Skill
	Playable          bool
}

/*
	races maps RaceType strings to a given Race struct.
*/
var races map[RaceType]*Race

/*
	init initializes the races map. It pairs each of the six default races with
	a fully populated Race struct.
*/
func init() {
	races = map[RaceType]*Race{
		Giant: &Race{
			Name:        string(Giant),
			Description: "A giant",
			Playable:    false,
			Attributes: &Attributes{
				AC: &AC{
					Magic: 100,
					Bash:  -100,
				},
				Stats: &Stats{
					Str:  19,
					Int:  12,
					Wis:  15,
					Dex:  14,
					Con:  18,
					Luck: 12,
				},
				Hit: 1,
				Dam: 2,
			},
		},
		Human: &Race{
			Name:        string(Human),
			Description: "A human",
			Playable:    true,
			Attributes: &Attributes{
				AC: &AC{},
				Stats: &Stats{
					Str:  15,
					Int:  15,
					Wis:  15,
					Dex:  15,
					Con:  15,
					Luck: 15,
				},
				Hit: 1,
				Dam: 1,
			},
		},
		Elf: &Race{
			Name:        string(Elf),
			Description: "An elf",
			Playable:    true,
			Attributes: &Attributes{
				AC: &AC{},
				Stats: &Stats{
					Str:  12,
					Int:  17,
					Wis:  18,
					Dex:  18,
					Con:  12,
					Luck: 13,
				},
				Hit: 1,
				Dam: 1,
			},
		},
		Dwarf: &Race{
			Name:        string(Dwarf),
			Description: "A dwarf",
			Playable:    true,
			Attributes: &Attributes{
				AC: &AC{
					Bash:   -100,
					Slash:  50,
					Pierce: 50,
				},
				Stats: &Stats{
					Str:  17,
					Int:  12,
					Wis:  16,
					Dex:  12,
					Con:  18,
					Luck: 15,
				},
				Hit: 1,
				Dam: 1,
			},
		},
		Gnome: &Race{
			Name:        string(Gnome),
			Description: "A gnome",
			Playable:    false,
			Attributes: &Attributes{
				AC: &AC{
					Magic: -100,
					Bash:  100,
				},
				Stats: &Stats{
					Str:  12,
					Int:  15,
					Wis:  19,
					Dex:  17,
					Con:  12,
					Luck: 15,
				},
				Hit: 1,
				Dam: 1,
			},
		},
		Nymph: &Race{
			Name:        string(Nymph),
			Description: "A nymph",
			Playable:    false,
			Attributes: &Attributes{
				AC: &AC{
					Magic: -100,
					Bash:  100,
				},
				Stats: &Stats{
					Str:  10,
					Int:  19,
					Wis:  19,
					Dex:  17,
					Con:  9,
					Luck: 16,
				},
				Hit: 1,
				Dam: 1,
			},
		},
	}
}
