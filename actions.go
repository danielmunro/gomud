package gomud

import "github.com/danielmunro/gomud/io"

var actions []*action

func newLookAction() *action {
	return &action{
		command:      io.LookCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition, sittingDisposition},
		mutator:      look,
		syntax:       []syntax{commandSyntax},
	}
}

func newKillAction() *action {
	return &action{
		command:      io.KillCommand,
		dispositions: []disposition{standingDisposition},
		mutator:      kill,
		syntax:       []syntax{commandSyntax, mobInRoomSyntax},
	}
}

func newFleeAction() *action {
	return &action{
		command:      io.FleeCommand,
		dispositions: []disposition{fightingDisposition},
		mutator:      flee,
		syntax:       []syntax{commandSyntax},
	}
}

func newWearAction() *action {
	return &action{
		command:      io.WearCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator:      wear,
		syntax:       []syntax{commandSyntax, itemInInventorySyntax},
	}
}

func newRemoveAction() *action {
	return &action{
		command:      io.RemoveCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator:      remove,
		syntax:       []syntax{commandSyntax, itemEquippedSyntax},
	}
}

func newGetAction() *action {
	return &action{
		command:      io.GetCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator:      get,
		syntax:       []syntax{commandSyntax, itemInRoomSyntax},
	}
}

func newDropAction() *action {
	return &action{
		command:      io.DropCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator:      drop,
		syntax:       []syntax{commandSyntax, itemInInventorySyntax},
	}
}

func newMoveAction(command io.Command, direction direction) *action {
	return &action{
		command: command,
		dispositions: []disposition{standingDisposition},
		mutator: func (b *io.Buffer, actionContext *ActionContext, eventService *EventService) *io.Output {
			return move(direction, b, actionContext, eventService)
		},
		chainToCommand: io.LookCommand,
	}
}

func newNoopAction() *action {
	return &action{
		command:      io.NoopCommand,
		dispositions: []disposition{},
		mutator: func(b *io.Buffer, actionContext *ActionContext, eventService *EventService) *io.Output {
			return io.NewOutputToRequestCreator(b, io.CompletedStatus, "What was that?")
		},
	}
}

func init() {
	actions = []*action{
		newLookAction(),
		newKillAction(),
		newFleeAction(),
		newWearAction(),
		newRemoveAction(),
		newGetAction(),
		newDropAction(),
		newMoveAction(io.NorthCommand, dNorth),
		newMoveAction(io.SouthCommand, dSouth),
		newMoveAction(io.EastCommand, dEast),
		newMoveAction(io.WestCommand, dWest),
		newMoveAction(io.UpCommand, dUp),
		newMoveAction(io.DownCommand, dDown),
	}
}

func findActionByCommand(command io.Command) *action {
	for _, a := range actions {
		if a.command == command {
			return a
		}
	}
	return newNoopAction()
}
