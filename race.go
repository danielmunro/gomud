package gomud

import (
	"github.com/jinzhu/gorm"
	"strings"
)

type RaceName string

const (
	ElfRace = "elf"
	DwarfRace = "dwarf"
	HumanRace = "human"
	CritterRace = "critter"
)

var races []*Race

func getRace(name RaceName) *Race {
	for _, i := range races {
		if strings.HasPrefix(string(i.Name), string(name)) {
			return i
		}
	}

	return races[len(races)]
}

type Race struct {
	gorm.Model
	Name       RaceName
	Attributes *Attributes
}

func NewRace(name RaceName, attrs *Attributes) *Race {
	return &Race{
		Name:       name,
		Attributes: attrs,
	}
}

func init() {
	races = []*Race{
		NewRace(ElfRace, 	 NewStats(12, 17, 16, 16, 12)),
		NewRace(DwarfRace,   NewStats(18, 12, 17, 11, 18)),
		NewRace(HumanRace,   NewStats(15, 15, 15, 15, 15)),
		NewRace(CritterRace, NewStats(15, 15, 15, 15, 15)),
	}
}
