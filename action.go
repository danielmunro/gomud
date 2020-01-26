package gomud

import (
	"fmt"
)

type mutator func(i *input, actionContext *ActionContext, eventService *EventService) *output

type action struct {
	command command
	dispositions []disposition
	mutator mutator
	syntax []syntax
	chainToCommand command
}

func (a *action) mobHasDisposition(mob *Mob) bool {
	for _, d := range a.dispositions {
		if d == mob.disposition {
			return true
		}
	}
	return false
}

func transferItem(i int, from []*item, to []*item) ([]*item, []*item) {
	item := from[i]
	from = append(from[0:i], from[i+1:]...)
	to = append(to, item)

	return from, to
}

func exitsString(r *room) string {
	var exits string

	for _, e := range r.exits {
		exits = fmt.Sprintf("%s%s", exits, string(e.direction[0]))
	}

	return fmt.Sprintf("[%s]", exits)
}

func mobsString(r *room, mob *Mob) string {
	var mobs string

	for _, m := range r.mobs {
		if m != mob {
			mobs = fmt.Sprintf("%s is here.\n%s", m.String(), mobs)
		}
	}

	return mobs
}

func itemsString(r *room) string {
	var items string

	for _, i := range r.items {
		items = fmt.Sprintf("%s is here.\n%s", i.String(), items)
	}

	return items
}
