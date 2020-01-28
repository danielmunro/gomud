package io

type Command string

const (
	LookCommand   Command = "look"
	NorthCommand  Command = "north"
	SouthCommand  Command = "south"
	EastCommand   Command = "east"
	WestCommand   Command = "west"
	UpCommand     Command = "up"
	DownCommand   Command = "down"
	GetCommand    Command = "get"
	DropCommand   Command = "drop"
	WearCommand   Command = "wear"
	RemoveCommand Command = "remove"
	KillCommand   Command = "kill"
	FleeCommand   Command = "flee"
	InventoryCommand Command = "inventory"
	NoopCommand   Command = "noop"
)

var Commands []Command

func init() {
	Commands = []Command{
		LookCommand,
		NorthCommand,
		SouthCommand,
		EastCommand,
		WestCommand,
		UpCommand,
		DownCommand,
		GetCommand,
		DropCommand,
		WearCommand,
		RemoveCommand,
		KillCommand,
		FleeCommand,
		InventoryCommand,
	}
}
