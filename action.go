package gomud

import (
	"github.com/danielmunro/gomud/io"
	"github.com/danielmunro/gomud/model"
)

type Mutator func(actionContext *ActionContext, actionService *ActionService) *io.Output

type Action struct {
	command        io.Command
	dispositions   []model.Disposition
	mutator        Mutator
	syntax         []syntax
	chainToCommand io.Command
}

func (a *Action) mobHasDisposition(mob *model.Mob) bool {
	for _, d := range a.dispositions {
		if mob.HasDisposition(d) {
			return true
		}
	}
	return false
}

