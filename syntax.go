package gomud

type syntax string

const (
	commandSyntax         syntax = "command"
	mobInRoomSyntax       syntax = "mob in room"
	itemInInventorySyntax syntax = "item in inventory"
	itemInRoomSyntax      syntax = "item in room"
	itemEquippedSyntax    syntax = "item equipped"
)
