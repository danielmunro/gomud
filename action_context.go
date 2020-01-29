package gomud

import "github.com/danielmunro/gomud/io"

type context struct {
	syntax syntax
	thing interface{}
	error error
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
	room *Room
	mob *Mob
	buffer *io.Buffer
	results []*context
}

func (ac *ActionContext) findErrorContext() *context {
	for _, r := range ac.results {
		if r.error != nil {
			return r
		}
	}
	return nil
}

func (ac *ActionContext) getMobBySyntax(syntax syntax) *Mob {
	for _, r := range ac.results {
		if r.syntax == syntax {
			return r.thing.(*Mob)
		}
	}
	return nil
}

func (ac *ActionContext) getItemBySyntax(syntax syntax) *Item {
	for _, r := range ac.results {
		if r.syntax == syntax {
			return r.thing.(*Item)
		}
	}
	return nil
}

func (ac *ActionContext) getExitBySyntax(syntax syntax) *exit {
	for _, r := range ac.results {
		if r.syntax == syntax {
			return r.thing.(*exit)
		}
	}
	return nil
}
