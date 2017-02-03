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
			&attributes{
				str: 12,
				int: 17,
				wis: 16,
				dex: 16,
				con: 12,
			},
		),
	)
	r = append(
		r,
		newRace(
			"dwarf",
			&attributes{
				str: 18,
				int: 12,
				wis: 17,
				dex: 11,
				con: 18,
			},
		),
	)
	r = append(
		r,
		newRace(
			"human",
			&attributes{
				str: 15,
				int: 15,
				wis: 15,
				dex: 15,
				con: 15,
			},
		),
	)
	r = append(
		r,
		newRace(
			"amorphous form",
			&attributes{
				str: 15,
				int: 15,
				wis: 15,
				dex: 15,
				con: 15,
			},
		),
	)
}
