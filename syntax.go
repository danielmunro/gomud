package gomud

type syntax string

const (
	// implemented
	commandSyntax               syntax = "command"
	mobInRoomSyntax             syntax = "mob in room"
	merchantInRoomSyntax        syntax = "merchant in room"
	itemInInventorySyntax       syntax = "item in inventory"
	itemInRoomSyntax            syntax = "item in room"
	itemEquippedSyntax          syntax = "item equipped"
	itemInTargetInventorySyntax syntax = "item in target inventory"

	// not implemented
	exitDirectionSyntax syntax = "exit direction"
)
