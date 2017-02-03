package gomud

import (
	"fmt"
	"strings"
)

type input struct {
	mob    *mob
	client *client
	args   []string
}

func newInput(c *client, args []string) *input {
	return &input{
		client: c,
		args:   args,
		mob:    c.mob,
	}
}

func (i *input) matchesCommand(s string) bool {
	return strings.HasPrefix(s, i.args[0])
}

func (i *input) matchesSubject(s []string) bool {
	for _, v := range s {
		if strings.HasPrefix(v, i.args[1]) {
			return true
		}
	}

	return false
}

func parse(i *input) {
	if i.matchesCommand("look") {
		look(i)
	} else if i.matchesCommand("north") {
		move(north, i)
	} else if i.matchesCommand("south") {
		move(south, i)
	} else if i.matchesCommand("east") {
		move(east, i)
	} else if i.matchesCommand("west") {
		move(west, i)
	} else if i.matchesCommand("up") {
		move(up, i)
	} else if i.matchesCommand("down") {
		move(down, i)
	} else if i.matchesCommand("get") {
		get(i)
	} else if i.matchesCommand("drop") {
		drop(i)
	} else if i.matchesCommand("wear") {
		wear(i)
	} else if i.matchesCommand("remove") {
		remove(i)
	} else {
		i.client.writePrompt("Eh?")
	}
}

func look(i *input) {
	r := i.client.mob.room
	i.client.writePrompt(
		fmt.Sprintf(
			"%s\n%s\n%s\n%s%s",
			r.name,
			r.description,
			exitsString(r),
			itemsString(r),
			mobsString(r, i.client.mob),
		),
	)
}

func wear(i *input) {
	for j, item := range i.mob.items {
		if i.matchesSubject(item.identifiers) {
			i.mob.items, i.mob.equipped = transferItem(j, i.mob.items, i.mob.equipped)
			i.mob.notify(fmt.Sprintf("You wear %s.", item.String()))
			return
		}
	}

	i.mob.notify("You can't find that.")
}

func remove(i *input) {
	for j, item := range i.mob.equipped {
		if i.matchesSubject(item.identifiers) {
			i.mob.equipped, i.mob.items = transferItem(j, i.mob.equipped, i.mob.items)
			i.mob.notify(fmt.Sprintf("You remove %s.", item.String()))
			return
		}
	}

	i.mob.notify("You can't find that.")
}

func get(i *input) {
	for j, item := range i.mob.room.items {
		if i.matchesSubject(item.identifiers) {
			i.mob.room.items, i.mob.items = transferItem(j, i.mob.room.items, i.mob.items)
			message := fmt.Sprintf("%s picks up %s.", i.mob.String(), item.String())
			for _, m := range i.mob.room.mobs {
				if m == i.mob {
					m.notify(fmt.Sprintf("You pick up %s.", item.String()))
				} else {
					m.notify(message)
				}
			}

			return
		}
	}
}

func drop(i *input) {
	for j, item := range i.mob.items {
		if i.matchesSubject(item.identifiers) {
			i.mob.items, i.mob.room.items = transferItem(j, i.mob.items, i.mob.room.items)
			message := fmt.Sprintf("%s drops %s.", i.mob.String(), item.String())
			for _, m := range i.mob.room.mobs {
				if m == i.mob {
					m.notify(fmt.Sprintf("You drop %s.", item.String()))
				} else {
					m.notify(message)
				}
			}

			return
		}
	}
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

func move(d direction, i *input) {
	for _, e := range i.client.mob.room.exits {
		if e.direction == d {
			i.client.mob.move(e)
			look(i)
			return
		}
	}
	i.client.writePrompt("Alas, you cannot go that way.")
}
