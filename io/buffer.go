package io

import (
	"fmt"
	"strings"
)

type Buffer struct {
	Client *Client
	Input  string
}

func NewBuffer(client *Client, input string) *Buffer {
	return &Buffer{
		client,
		input,
	}
}

func (b *Buffer) GetCommand() Command {
	args := strings.Split(b.Input, " ")
	for _, c := range Commands {
		if isCommand(c, args[0]) == true {
			return c
		}
	}

	return NoopCommand
}

func (b *Buffer) MatchesSubject(s []string) bool {
	args := strings.Split(b.Input, " ")
	for _, v := range s {
		if strings.HasPrefix(v, args[1]) {
			return true
		}
	}

	return false
}

func (b *Buffer) ToString() string {
	return fmt.Sprintf("Client: %s, Input: '%s'", b.Client.id, b.Input)
}

func (b *Buffer) CreateOutputToRequestCreator(messageToRequestCreator string) *Output {
	return NewOutputToRequestCreator(b, CompletedStatus, messageToRequestCreator)
}

func (b *Buffer) CreateOutput(messageToRequestCreator string, messageToTarget string, messageToObservers string) *Output {
	return NewOutput(b, CompletedStatus, messageToRequestCreator, messageToTarget, messageToObservers)
}

func isCommand(c Command, p string) bool {
	return strings.HasPrefix(string(c), p)
}
