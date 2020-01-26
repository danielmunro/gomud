package gomud

import "fmt"

type Buffer struct {
	client *client
	input string
}

func NewBuffer(client *client, input string) *Buffer {
	return &Buffer{
		client,
		input,
	}
}

func (b *Buffer) ToString() string {
	return fmt.Sprintf("client: %s, input: '%s'", b.client.id, b.input)
}
