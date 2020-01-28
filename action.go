package gomud

import (
	"fmt"
	"github.com/danielmunro/gomud/io"
)

type Mutator func(actionContext *ActionContext, actionService *ActionService) *io.Output

type Action struct {
	command        io.Command
	dispositions   []disposition
	mutator        Mutator
	syntax         []syntax
	chainToCommand io.Command
}

func (a *Action) mobHasDisposition(mob *Mob) bool {
	for _, d := range a.dispositions {
		if d == mob.disposition {
			return true
		}
	}
	return false
}

func transferItem(item *item, from []*item, to []*item) ([]*item, []*item) {
	for i, x := range from {
		if x == item {
			from = append(from[0:i], from[i+1:]...)
			to = append(to, item)
		}
	}

	return from, to
}

func transferItemByIndex(i int, from []*item, to []*item) ([]*item, []*item) {
	item := from[i]
	from = append(from[0:i], from[i+1:]...)
	to = append(to, item)

	return from, to
}

func exitsString(r *Room) string {
	var exits string

	for _, e := range r.exits {
		exits = fmt.Sprintf("%s%s", exits, string(e.direction[0]))
	}

	return fmt.Sprintf("[%s]", exits)
}

func mobsString(mobs []*Mob) string {
	var buf string

	for _, m := range mobs {
		buf = fmt.Sprintf("%s is here.\n%s", m.String(), buf)
	}

	return buf
}

func itemsString(r *Room) string {
	var items string

	for _, i := range r.items {
		items = fmt.Sprintf("%s is here.\n%s", i.String(), items)
	}

	return items
}
