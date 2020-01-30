package gomud

import (
	"github.com/danielmunro/gomud/io"
	"github.com/danielmunro/gomud/model"
)

type Login struct {
	client *io.Client
	mob    *model.Mob
}

func NewLogin(client *io.Client, mob *model.Mob) *Login {
	return &Login{
		client,
		mob,
	}
}
