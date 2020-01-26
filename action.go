package gomud

import (
	"fmt"
)

type mutator func(i *input, actionContext *ActionContext) *output

type action struct {
	command command
	dispositions []disposition
	mutator mutator
	syntax []syntax
}

func (a *action) mobHasDisposition(mob *mob) bool {
	for _, d := range a.dispositions {
		if d == mob.disposition {
			return true
		}
	}
	return false
}

func kill(i *input, actionContext *ActionContext) *output {
	mob := actionContext.getMobBySyntax(mobInRoomSyntax)
	newFight(i.mob, mob)
	return newOutputToRequestCreator(i, CompletedStatus, "You scream and attack!")
}

func flee(i *input, actionContext *ActionContext) *output {
	i.mob.fight = nil
	i.mob.move(i.mob.room.exits[dice().Intn(len(i.mob.room.exits))])
	return newOutputToRequestCreator(i, CompletedStatus, "you flee!")
}

func look(i *input, actionContext *ActionContext) *output {
	r := i.room
	return newOutputToRequestCreator(
		i,
		CompletedStatus,
		fmt.Sprintf(
			"%s\n%s\n%s\n%s%s",
			r.name,
			r.description,
			exitsString(r),
			itemsString(r),
			mobsString(r, i.mob),
		),
	)
}

func wear(i *input, actionContext *ActionContext) *output {
	for j, item := range i.mob.items {
		if i.matchesSubject(item.identifiers) {
			for k, eq := range i.mob.equipped {
				if eq.position == item.position {
					i.mob.equipped, i.mob.items = transferItem(k, i.mob.equipped, i.mob.items)
					i.mob.notify(fmt.Sprintf("You remove %s and put it in your inventory.", eq.String()))
				}
			}
			i.mob.items, i.mob.equipped = transferItem(j, i.mob.items, i.mob.equipped)
			return newOutputToRequestCreator(i, CompletedStatus, fmt.Sprintf("You wear %s.", item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func remove(i *input, actionContext *ActionContext) *output {
	for j, item := range i.mob.equipped {
		if i.matchesSubject(item.identifiers) {
			i.mob.equipped, i.mob.items = transferItem(j, i.mob.equipped, i.mob.items)
			return newOutputToRequestCreator(i, CompletedStatus, fmt.Sprintf("You remove %s.", item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func get(i *input, actionContext *ActionContext) *output {
	for j, item := range i.room.items {
		if i.matchesSubject(item.identifiers) {
			i.room.items, i.mob.items = transferItem(j, i.room.items, i.mob.items)
			return newOutput(
				i,
				CompletedStatus,
				fmt.Sprintf("You pick up %s", item.String()),
				fmt.Sprintf("%s picks up %s", i.mob.name, item.String()),
				fmt.Sprintf("%s picks up %s", i.mob.name, item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func drop(i *input, actionContext *ActionContext) *output {
	for j, item := range i.mob.items {
		if i.matchesSubject(item.identifiers) {
			i.mob.items, i.room.items = transferItem(j, i.mob.items, i.room.items)
			return newOutput(
				i,
				CompletedStatus,
				fmt.Sprintf("You drop %s", item.String()),
				fmt.Sprintf("%s drops %s", i.mob.name, item.String()),
				fmt.Sprintf("%s drops %s", i.mob.name, item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func move(d direction, i *input, actionContext *ActionContext) *output {
	for _, e := range i.room.exits {
		if e.direction == d {
			i.client.mob.move(e)
		}
	}
	return newOutputToRequestCreator(i, FailedStatus, "Alas, you cannot go that way.")
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

func mobsString(r *room, mob *mob) string {
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
