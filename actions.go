package gomud

import (
	"github.com/danielmunro/gomud/io"
	"github.com/danielmunro/gomud/message"
	"github.com/danielmunro/gomud/model"
)

var actions []*Action

func newLookAction() *Action {
	return &Action{
		command: io.LookCommand,
		dispositions: []model.Disposition{
			model.StandingDisposition,
			model.FightingDisposition,
			model.SittingDisposition,
		},
		mutator: look,
		syntax:  []syntax{commandSyntax},
	}
}

func newKillAction() *Action {
	return &Action{
		command:      io.KillCommand,
		dispositions: []model.Disposition{model.StandingDisposition},
		mutator:      kill,
		syntax:       []syntax{commandSyntax, mobInRoomSyntax},
	}
}

func newFleeAction() *Action {
	return &Action{
		command:      io.FleeCommand,
		dispositions: []model.Disposition{model.FightingDisposition},
		mutator:      flee,
		syntax:       []syntax{commandSyntax},
	}
}

func newWearAction() *Action {
	return &Action{
		command: io.WearCommand,
		dispositions: []model.Disposition{
			model.StandingDisposition,
			model.FightingDisposition,
		},
		mutator: wear,
		syntax:  []syntax{commandSyntax, itemInInventorySyntax},
	}
}

func newRemoveAction() *Action {
	return &Action{
		command: io.RemoveCommand,
		dispositions: []model.Disposition{
			model.StandingDisposition,
			model.FightingDisposition,
		},
		mutator: remove,
		syntax:  []syntax{commandSyntax, itemEquippedSyntax},
	}
}

func newGetAction() *Action {
	return &Action{
		command: io.GetCommand,
		dispositions: []model.Disposition{
			model.StandingDisposition,
			model.FightingDisposition,
		},
		mutator: get,
		syntax:  []syntax{commandSyntax, itemInRoomSyntax},
	}
}

func newDropAction() *Action {
	return &Action{
		command: io.DropCommand,
		dispositions: []model.Disposition{
			model.StandingDisposition,
			model.FightingDisposition,
		},
		mutator: drop,
		syntax:  []syntax{commandSyntax, itemInInventorySyntax},
	}
}

func newInventoryAction() *Action {
	return &Action{
		command: io.InventoryCommand,
		dispositions: []model.Disposition{
			model.StandingDisposition,
			model.FightingDisposition,
			model.SittingDisposition,
			model.SleepingDisposition,
		},
		mutator: inventory,
		syntax:  []syntax{commandSyntax},
	}
}

func newSitAction() *Action {
	return &Action{
		command: io.SitCommand,
		dispositions: []model.Disposition{
			model.StandingDisposition,
			model.SleepingDisposition,
		},
		mutator: sit,
		syntax:  []syntax{commandSyntax},
	}
}

func newSleepAction() *Action {
	return &Action{
		command: io.SleepCommand,
		dispositions: []model.Disposition{
			model.StandingDisposition,
			model.SittingDisposition,
		},
		mutator: sleep,
		syntax:  []syntax{commandSyntax},
	}
}

func newWakeAction() *Action {
	return &Action{
		command: io.WakeCommand,
		dispositions: []model.Disposition{
			model.SittingDisposition,
			model.SleepingDisposition,
		},
		mutator: wake,
		syntax:  []syntax{commandSyntax},
	}
}

func newMoveAction(command io.Command, direction model.Direction) *Action {
	return &Action{
		command:      command,
		dispositions: []model.Disposition{model.StandingDisposition},
		mutator: func(actionContext *ActionContext, actionService *ActionService) *io.Output {
			return move(direction, actionContext, actionService)
		},
		syntax:         []syntax{exitDirectionSyntax},
		chainToCommand: io.LookCommand,
	}
}

func newListAction() *Action {
	return &Action{
		command:      io.ListCommand,
		dispositions: []model.Disposition{model.StandingDisposition},
		mutator:      list,
		syntax:       []syntax{merchantInRoomSyntax},
	}
}

func newNoopAction() *Action {
	return &Action{
		command:      io.NoopCommand,
		dispositions: []model.Disposition{},
		mutator: func(actionContext *ActionContext, actionService *ActionService) *io.Output {
			return actionContext.buffer.CreateOutputToRequestCreator(message.ErrorInputNotUnderstood)
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
		newSitAction(),
		newSleepAction(),
		newWakeAction(),
		newListAction(),
		newMoveAction(io.NorthCommand, model.NorthDirection),
		newMoveAction(io.SouthCommand, model.SouthDirection),
		newMoveAction(io.EastCommand, model.EastDirection),
		newMoveAction(io.WestCommand, model.WestDirection),
		newMoveAction(io.UpCommand, model.UpDirection),
		newMoveAction(io.DownCommand, model.DownDirection),
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
