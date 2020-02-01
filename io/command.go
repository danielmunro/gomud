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
	SitCommand Command = "sit"
	WakeCommand Command = "wake"
	SleepCommand Command = "sleep"
	ListCommand Command = "list"
	SellCommand Command = "sell"
	BuyCommand Command = "buy"
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
		SitCommand,
		WakeCommand,
		SleepCommand,
		ListCommand,
		SellCommand,
		BuyCommand,
	}
}
