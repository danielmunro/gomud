package gomud

import (
	"github.com/danielmunro/gomud/io"
)

type Mutator func(actionContext *ActionContext, actionService *ActionService) *io.Output

type Action struct {
	command        io.Command
	dispositions   []disposition
	mutator        Mutator
	syntax         []syntax
	chainToCommand io.Command
}

func (a *Action) mobHasDisposition(mob *Mob) bool {
	for _, d := range a.dispositions {
		if d == mob.disposition {
			return true
		}
	}
	return false
}

