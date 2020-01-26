package gomud

type command string

const (
	LookCommand   command = "look"
	NorthCommand  command = "north"
	SouthCommand  command = "south"
	EastCommand   command = "east"
	WestCommand   command = "west"
	UpCommand     command = "up"
	DownCommand   command = "down"
	GetCommand    command = "get"
	DropCommand   command = "drop"
	WearCommand   command = "wear"
	RemoveCommand command = "remove"
	KillCommand   command = "kill"
	FleeCommand   command = "flee"
	NoopCommand   command = "noop"
)

var commands []command

func init() {
	commands = []command{LookCommand, NorthCommand, SouthCommand, EastCommand, WestCommand, UpCommand, DownCommand, GetCommand, DropCommand, WearCommand, RemoveCommand, KillCommand, FleeCommand}
}
