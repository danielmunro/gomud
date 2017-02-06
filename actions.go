package gomud

import (
	"fmt"
	"strings"
)

type action struct {
	i *input
}

func newAction(m *mob, i string) {
	newActionWithInput(&input{mob: m, args: strings.Split(i, " ")})
}

func newActionWithInput(i *input) {
	a := &action{
		i: i,
	}

	switch a.i.getCommand() {
	case cLook:
		a.look()
		return
	case cNorth:
		a.move(dNorth)
		return
	case cSouth:
		a.move(dSouth)
		return
	case cEast:
		a.move(dEast)
		return
	case cWest:
		a.move(dWest)
		return
	case cUp:
		a.move(dUp)
		return
	case cDown:
		a.move(dDown)
		return
	case cDrop:
		a.drop()
		return
	case cGet:
		a.get()
		return
	case cWear:
		a.wear()
		return
	case cRemove:
		a.remove()
		return
	default:
		i.client.writePrompt("Eh?")
	}

}

func (a *action) look() {
	r := a.i.mob.room
	a.i.mob.notify(
		fmt.Sprintf(
			"%s\n%s\n%s\n%s%s",
			r.name,
			r.description,
			exitsString(r),
			itemsString(r),
			mobsString(r, a.i.mob),
		),
	)
}

func (a *action) wear() {
	for j, item := range a.i.mob.items {
		if a.i.matchesSubject(item.identifiers) {
			for k, eq := range a.i.mob.equipped {
				if eq.position == item.position {
					a.i.mob.equipped, a.i.mob.items = transferItem(k, a.i.mob.equipped, a.i.mob.items)
					a.i.mob.notify(fmt.Sprintf("You remove %s and put it in your inventory.", eq.String()))
				}
			}
			a.i.mob.items, a.i.mob.equipped = transferItem(j, a.i.mob.items, a.i.mob.equipped)
			a.i.mob.notify(fmt.Sprintf("You wear %s.", item.String()))
			return
		}
	}

	a.i.mob.notify("You can't find that.")
}

func (a *action) remove() {
	for j, item := range a.i.mob.equipped {
		if a.i.matchesSubject(item.identifiers) {
			a.i.mob.equipped, a.i.mob.items = transferItem(j, a.i.mob.equipped, a.i.mob.items)
			a.i.mob.notify(fmt.Sprintf("You remove %s.", item.String()))
			return
		}
	}

	a.i.mob.notify("You can't find that.")
}

func (a *action) get() {
	for j, item := range a.i.mob.room.items {
		if a.i.matchesSubject(item.identifiers) {
			a.i.mob.room.items, a.i.mob.items = transferItem(j, a.i.mob.room.items, a.i.mob.items)
			message := fmt.Sprintf("%s picks up %s.", a.i.mob.String(), item.String())
			for _, m := range a.i.mob.room.mobs {
				if m == a.i.mob {
					m.notify(fmt.Sprintf("You pick up %s.", item.String()))
				} else {
					m.notify(message)
				}
			}

			return
		}
	}
}

func (a *action) drop() {
	for j, item := range a.i.mob.items {
		if a.i.matchesSubject(item.identifiers) {
			a.i.mob.items, a.i.mob.room.items = transferItem(j, a.i.mob.items, a.i.mob.room.items)
			message := fmt.Sprintf("%s drops %s.", a.i.mob.String(), item.String())
			for _, m := range a.i.mob.room.mobs {
				if m == a.i.mob {
					m.notify(fmt.Sprintf("You drop %s.", item.String()))
				} else {
					m.notify(message)
				}
			}

			return
		}
	}
}

func (a *action) move(d direction) {
	for _, e := range a.i.client.mob.room.exits {
		if e.direction == d {
			a.i.client.mob.move(e)
			newAction(a.i.client.mob, "look")
			return
		}
	}
	a.i.client.writePrompt("Alas, you cannot go that way.")
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
