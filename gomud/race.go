package gomud

type RaceType string

const (
	Giant RaceType = "giant"
	Human RaceType = "human"
)

type Race struct {
	Name, Description string
	Attributes        *Attributes
	Skills            []*Skill
}

var races map[RaceType]*Race

func init() {
	races = map[RaceType]*Race{
		Giant: &Race{
			Name:        string(Giant),
			Description: "A giant",
		},
		Human: &Race{
			Name:        string(Human),
			Description: "A human",
		},
	}
}
