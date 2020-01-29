package gomud

import "strings"

type RaceName string

const (
	ElfRace = "elf"
	DwarfRace = "dwarf"
	HumanRace = "human"
	CritterRace = "critter"
)

var races []*Race

func getRace(n string) *Race {
	for _, i := range races {
		if strings.HasPrefix(string(i.Name), n) {
			return i
		}
	}

	return races[len(races)]
}

type Race struct {
	Name       RaceName
	Attributes *attributes
}

func NewRace(name RaceName, attrs *attributes) *Race {
	return &Race{
		Name:       name,
		Attributes: attrs,
	}
}

func init() {
	races = []*Race{
		NewRace(ElfRace, 	 newStats(12, 17, 16, 16, 12)),
		NewRace(DwarfRace,   newStats(18, 12, 17, 11, 18)),
		NewRace(HumanRace,   newStats(15, 15, 15, 15, 15)),
		NewRace(CritterRace, newStats(15, 15, 15, 15, 15)),
	}
}
