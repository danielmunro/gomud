package gomud

import "strings"

var r []*race

func getRace(n string) *race {
	for _, i := range r {
		if strings.HasPrefix(i.name, n) {
			return i
		}
	}

	return r[len(r)]
}

type race struct {
	name  string
	attrs *attributes
}

func newRace(name string, attrs *attributes) *race {
	return &race{
		name:  name,
		attrs: attrs,
	}
}

func init() {
	r = append(
		r,
		newRace(
			"elf",
			newAttributes(map[attribute]int{aStr: 12, aInt: 17, aWis: 16, aDex: 16, aCon: 12}),
		),
	)
	r = append(
		r,
		newRace(
			"dwarf",
			newAttributes(map[attribute]int{aStr: 18, aInt: 12, aWis: 17, aDex: 11, aCon: 18}),
		),
	)
	r = append(
		r,
		newRace(
			"human",
			newAttributes(map[attribute]int{aStr: 15, aInt: 15, aWis: 15, aDex: 15, aCon: 15}),
		),
	)
	r = append(
		r,
		newRace(
			"critter",
			newAttributes(map[attribute]int{aStr: 15, aInt: 15, aWis: 15, aDex: 15, aCon: 15}),
		),
	)
}
