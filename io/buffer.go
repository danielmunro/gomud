package io

import (
	"fmt"
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

func (b *Buffer) ToString() string {
	return fmt.Sprintf("Client: %s, Input: '%s'", b.Client.id, b.Input)
}
