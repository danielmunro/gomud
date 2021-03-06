package gomud

import (
	"github.com/danielmunro/gomud/io"
	"github.com/danielmunro/gomud/message"
	"github.com/danielmunro/gomud/model"
	"reflect"
)

type context struct {
	syntax syntax
	thing  interface{}
	error  error
}

func newContext(syntax syntax, thing interface{}, error error) *context {
	return &context{
		syntax,
		thing,
		error,
	}
}

type ActionContext struct {
	hasDisposition bool
	room           *model.Room
	mob            *model.Mob
	buffer         *io.Buffer
	results        []*context
}

func (ac *ActionContext) CreateOutput(message *message.Message) *io.Output {
	return ac.buffer.CreateOutput(message.ToRequestCreator, message.ToTarget, message.ToObservers)
}

func (ac *ActionContext) findErrorContext() *context {
	for _, r := range ac.results {
		if r.error != nil {
			return r
		}
	}
	return nil
}

func (ac *ActionContext) getMobBySyntax(syntax syntax) *model.Mob {
	for _, r := range ac.results {
		if r.syntax == syntax {
			return r.thing.(*model.Mob)
		}
	}
	return nil
}

func (ac *ActionContext) getItemBySyntax(syntax syntax) *model.Item {
	for _, r := range ac.results {
		if r.syntax == syntax {
			return r.thing.(*model.Item)
		}
	}
	return nil
}

func (ac *ActionContext) getExitBySyntax(syntax syntax) *model.Exit {
	for _, r := range ac.results {
		if r.syntax == syntax {
			return r.thing.(*model.Exit)
		}
	}
	return nil
}

func (ac *ActionContext) getFirstMob() *model.Mob {
	for _, r := range ac.results {
		if reflect.TypeOf(r.thing).String() == "*model.Mob" {
			return r.thing.(*model.Mob)
		}
	}
	return nil
}
