package gomud

/*
	Affect represents a Mob's behavior while performing a task.
	Name - the name of the affect.
	Duration - how long the affect lasts, in server ticks.
	Attributes - an Attributes representing how the affect modifies a
				Mob's attributes for the duration.
*/
type Affect struct {
	Name       string
	Duration   int
	Attributes *Attributes
}
