package gomud

type syntax string

const (
	commandSyntax         syntax = "Command"
	mobInRoomSyntax       syntax = "Mob in Room"
	itemInInventorySyntax syntax = "item in inventory"
	itemInRoomSyntax      syntax = "item in Room"
	itemEquippedSyntax    syntax = "item equipped"
)
