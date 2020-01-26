package gomud

import (
	"fmt"
	"log"
)

func kill(i *input, actionContext *ActionContext, eventService *EventService) *output {
	mob := actionContext.getMobBySyntax(mobInRoomSyntax)
	newFight(i.mob, mob)
	return newOutputToRequestCreator(i, CompletedStatus, "You scream and attack!")
}

func flee(i *input, actionContext *ActionContext, eventService *EventService) *output {
	i.mob.fight = nil
	i.mob.move(i.mob.room.exits[dice().Intn(len(i.mob.room.exits))])
	return newOutputToRequestCreator(i, CompletedStatus, "you flee!")
}

func look(i *input, actionContext *ActionContext, eventService *EventService) *output {
	r := i.room
	return newOutputToRequestCreator(
		i,
		CompletedStatus,
		fmt.Sprintf(
			"%s\n%s\n%s\n%s%s",
			r.name,
			r.description,
			exitsString(r),
			itemsString(r),
			mobsString(r, i.mob),
		),
	)
}

func wear(i *input, actionContext *ActionContext, eventService *EventService) *output {
	for j, item := range i.mob.items {
		if i.matchesSubject(item.identifiers) {
			for k, eq := range i.mob.equipped {
				if eq.position == item.position {
					i.mob.equipped, i.mob.items = transferItem(k, i.mob.equipped, i.mob.items)
					//i.mob.notify(fmt.Sprintf("You remove %s and put it in your inventory.", eq.String()))
				}
			}
			i.mob.items, i.mob.equipped = transferItem(j, i.mob.items, i.mob.equipped)
			return newOutputToRequestCreator(i, CompletedStatus, fmt.Sprintf("You wear %s.", item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func remove(i *input, actionContext *ActionContext, eventService *EventService) *output {
	for j, item := range i.mob.equipped {
		if i.matchesSubject(item.identifiers) {
			i.mob.equipped, i.mob.items = transferItem(j, i.mob.equipped, i.mob.items)
			return newOutputToRequestCreator(i, CompletedStatus, fmt.Sprintf("You remove %s.", item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func get(i *input, actionContext *ActionContext, eventService *EventService) *output {
	for j, item := range i.room.items {
		if i.matchesSubject(item.identifiers) {
			i.room.items, i.mob.items = transferItem(j, i.room.items, i.mob.items)
			return newOutput(
				i,
				CompletedStatus,
				fmt.Sprintf("You pick up %s", item.String()),
				fmt.Sprintf("%s picks up %s", i.mob.name, item.String()),
				fmt.Sprintf("%s picks up %s", i.mob.name, item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func drop(i *input, actionContext *ActionContext, eventService *EventService) *output {
	for j, item := range i.mob.items {
		if i.matchesSubject(item.identifiers) {
			i.mob.items, i.room.items = transferItem(j, i.mob.items, i.room.items)
			return newOutput(
				i,
				CompletedStatus,
				fmt.Sprintf("You drop %s", item.String()),
				fmt.Sprintf("%s drops %s", i.mob.name, item.String()),
				fmt.Sprintf("%s drops %s", i.mob.name, item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func move(d direction, i *input, actionContext *ActionContext, eventService *EventService) *output {
	log.Printf("move: direction: %s, command: %s", d, i.getCommand())
	for _, e := range i.room.exits {
		if e.direction == d {
			eventService.Publish(&Event{
				eventType:MobMoveEventType,
				mob: i.mob,
				room: e.room,
			})
			return newOutputToRequestCreator(i, CompletedStatus, fmt.Sprintf("You move %s.", d))
		}
	}
	return newOutputToRequestCreator(i, FailedStatus, "Alas, you cannot go that way.")
}
