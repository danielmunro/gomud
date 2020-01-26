package gomud

import "github.com/danielmunro/gomud/io"

type Login struct {
	client *io.Client
	mob *Mob
}

func NewLogin(client *io.Client, mob *Mob) *Login {
	return &Login{
		client,
		mob,
	}
}
