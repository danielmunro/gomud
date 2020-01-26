package gomud

type syntax string

const (
	commandSyntax         syntax = "command"
	mobInRoomSyntax       syntax = "Mob in room"
	itemInInventorySyntax syntax = "item in inventory"
	itemInRoomSyntax      syntax = "item in room"
	itemEquippedSyntax    syntax = "item equipped"
)
