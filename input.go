package gomud

import "strings"

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

func (i *input) getCommand() command {
	for _, c := range commands {
		if isCommand(c, i.args[0]) == true {
			return c
		}
	}

	return cNoop
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
