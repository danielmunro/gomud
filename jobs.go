package gomud

import "strings"

var j []*job

func getJob(n string) *job {
	for _, i := range j {
		if strings.HasPrefix(i.name, n) {
			return i
		}
	}

	return j[len(j)]
}

type job struct {
	name  string
	attrs *attributes
}

func newJob(name string, attrs *attributes) *job {
	return &job{
		name:  name,
		attrs: attrs,
	}
}

func init() {
	j = append(
		j,
		newJob(
			"mage",
			&attributes{
				int: 2,
				wis: 1,
				dex: 1,
			},
		),
	)
	j = append(
		j,
		newJob(
			"cleric",
			&attributes{
				wis: 2,
				int: 1,
				con: 1,
			},
		),
	)
	j = append(
		j,
		newJob(
			"thief",
			&attributes{
				dex: 2,
				con: 1,
				str: 1,
			},
		),
	)
	j = append(
		j,
		newJob(
			"warrior",
			&attributes{
				str: 2,
				con: 1,
				dex: 1,
			},
		),
	)
}
