package gomud

type context struct {
	syntax syntax
	thing interface{}
}

type ActionContext struct {
	hasDisposition bool
	results []*context
}

func (ac *ActionContext) getMobBySyntax(syntax syntax) *Mob {
	for _, r := range ac.results {
		if r.syntax == syntax {
			return r.thing.(*Mob)
		}
	}
	return nil
}
