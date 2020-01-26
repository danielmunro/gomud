package io

import (
	"strings"
)

type Input struct {
	client *Client
	args   []string
}

func NewInput(client *Client, args []string) *Input {
	return &Input{
		client,
		args,
	}
}

func (i *Input) GetCommand() Command {
	for _, c := range Commands {
		if isCommand(c, i.args[0]) == true {
			return c
		}
	}

	return NoopCommand
}

func (i *Input) MatchesSubject(s []string) bool {
	for _, v := range s {
		if strings.HasPrefix(v, i.args[1]) {
			return true
		}
	}

	return false
}

func isCommand(c Command, p string) bool {
	return strings.HasPrefix(string(c), p)
}
