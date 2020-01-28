package gomud

import "github.com/danielmunro/gomud/io"

var actions []*Action

func newLookAction() *Action {
	return &Action{
		command:      io.LookCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition, sittingDisposition},
		mutator:      look,
		syntax:       []syntax{commandSyntax},
	}
}

func newKillAction() *Action {
	return &Action{
		command:      io.KillCommand,
		dispositions: []disposition{standingDisposition},
		mutator:      kill,
		syntax:       []syntax{commandSyntax, mobInRoomSyntax},
	}
}

func newFleeAction() *Action {
	return &Action{
		command:      io.FleeCommand,
		dispositions: []disposition{fightingDisposition},
		mutator:      flee,
		syntax:       []syntax{commandSyntax},
	}
}

func newWearAction() *Action {
	return &Action{
		command:      io.WearCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator:      wear,
		syntax:       []syntax{commandSyntax, itemInInventorySyntax},
	}
}

func newRemoveAction() *Action {
	return &Action{
		command:      io.RemoveCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator:      remove,
		syntax:       []syntax{commandSyntax, itemEquippedSyntax},
	}
}

func newGetAction() *Action {
	return &Action{
		command:      io.GetCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator:      get,
		syntax:       []syntax{commandSyntax, itemInRoomSyntax},
	}
}

func newDropAction() *Action {
	return &Action{
		command:      io.DropCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator:      drop,
		syntax:       []syntax{commandSyntax, itemInInventorySyntax},
	}
}

func newInventoryAction() *Action {
	return &Action{
		command:      io.InventoryCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition, sittingDisposition, sleepingDisposition},
		mutator:      inventory,
		syntax:       []syntax{commandSyntax},
	}
}

func newMoveAction(command io.Command, direction direction) *Action {
	return &Action{
		command: command,
		dispositions: []disposition{standingDisposition},
		mutator: func (actionContext *ActionContext, actionService *ActionService) *io.Output {
			return move(direction, actionContext, actionService)
		},
		syntax: []syntax{exitDirectionSyntax},
		chainToCommand: io.LookCommand,
	}
}

func newNoopAction() *Action {
	return &Action{
		command:      io.NoopCommand,
		dispositions: []disposition{},
		mutator: func(actionContext *ActionContext, actionService *ActionService) *io.Output {
			return actionContext.buffer.CreateOutputToRequestCreator("What was that?")
		},
	}
}

func init() {
	actions = []*Action{
		newLookAction(),
		newKillAction(),
		newFleeAction(),
		newWearAction(),
		newRemoveAction(),
		newGetAction(),
		newDropAction(),
		newInventoryAction(),
		newMoveAction(io.NorthCommand, dNorth),
		newMoveAction(io.SouthCommand, dSouth),
		newMoveAction(io.EastCommand, dEast),
		newMoveAction(io.WestCommand, dWest),
		newMoveAction(io.UpCommand, dUp),
		newMoveAction(io.DownCommand, dDown),
	}
}

func findActionByCommand(command io.Command) *Action {
	for _, a := range actions {
		if a.command == command {
			return a
		}
	}
	return newNoopAction()
}
