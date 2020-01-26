package gomud

import (
	"github.com/danielmunro/gomud/io"
	"strings"
)

type input struct {
	mob    *Mob
	client *io.Client
	room   *room
	args   []string
}

func newInput(client *io.Client, mob *Mob, room *room, args []string) *input {
	return &input{
		mob,
		client,
		room,
		args,
	}
}

func (i *input) getCommand() command {
	for _, c := range commands {
		if isCommand(c, i.args[0]) == true {
			return c
		}
	}

	return NoopCommand
}

func (i *input) matchesSubject(s []string) bool {
	for _, v := range s {
		if strings.HasPrefix(v, i.args[1]) {
			return true
		}
	}

	return false
}

func isCommand(c command, p string) bool {
	return strings.HasPrefix(string(c), p)
}
