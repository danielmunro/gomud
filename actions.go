package gomud

var actions []*action

func newLookAction() *action {
	return &action{
		command: LookCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition, sittingDisposition},
		mutator: look,
		syntax: []syntax{commandSyntax},
	}
}

func newKillAction() *action {
	return &action{
		command: KillCommand,
		dispositions: []disposition{standingDisposition},
		mutator: kill,
		syntax: []syntax{commandSyntax, mobInRoomSyntax},
	}
}

func newFleeAction() *action {
	return &action{
		command: FleeCommand,
		dispositions: []disposition{fightingDisposition},
		mutator: flee,
		syntax: []syntax{commandSyntax},
	}
}

func newWearAction() *action {
	return &action{
		command: WearCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator: wear,
		syntax: []syntax{commandSyntax, itemInInventorySyntax},
	}
}

func newRemoveAction() *action {
	return &action{
		command: RemoveCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator: remove,
		syntax: []syntax{commandSyntax, itemEquippedSyntax},
	}
}

func newGetAction() *action {
	return &action{
		command: GetCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator: get,
		syntax: []syntax{commandSyntax, itemInRoomSyntax},
	}
}

func newDropAction() *action {
	return &action{
		command: DropCommand,
		dispositions: []disposition{standingDisposition, fightingDisposition},
		mutator: drop,
		syntax: []syntax{commandSyntax, itemInInventorySyntax},
	}
}

func newMoveAction(command command, direction direction) *action {
	return &action{
		command: command,
		dispositions: []disposition{standingDisposition},
		mutator: func (i *input, actionContext *ActionContext, eventService *EventService) *output {
			return move(direction, i, actionContext, eventService)
		},
		chainToCommand: LookCommand,
	}
}

func newNoopAction() *action {
	return &action{
		command: NoopCommand,
		dispositions: []disposition{},
		mutator: func(i *input, actionContext *ActionContext, eventService *EventService) *output {
			return newOutputToRequestCreator(i, CompletedStatus, "What was that?")
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
		newMoveAction(NorthCommand, dNorth),
		newMoveAction(SouthCommand, dSouth),
		newMoveAction(EastCommand, dEast),
		newMoveAction(WestCommand, dWest),
		newMoveAction(UpCommand, dUp),
		newMoveAction(DownCommand, dDown),
	}
}

func findActionByCommand(command command) *action {
	for _, a := range actions {
		if a.command == command {
			return a
		}
	}
	return newNoopAction()
}
