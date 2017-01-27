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
	}
}

func (i *input) matches(command string) bool {
	return strings.HasPrefix(command, i.args[0])
}

func parse(i *input) {
	if i.matches("look") {
		look(i)
	} else if i.matches("north") {
		move(north, i)
	} else if i.matches("south") {
		move(south, i)
	} else if i.matches("east") {
		move(east, i)
	} else if i.matches("west") {
		move(west, i)
	} else if i.matches("up") {
		move(up, i)
	} else if i.matches("down") {
		move(down, i)
	} else {
		i.client.writePrompt("Eh?")
	}
}

func look(i *input) {
	r := i.client.mob.room
	i.client.writePrompt(fmt.Sprintf("%s\n%s\n%s\n%s", r.name, r.description, r.exitsString(), r.mobsString(i.client.mob)))
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
