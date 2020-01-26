package gomud

import (
	"fmt"
	"github.com/danielmunro/gomud/io"
	"log"
)

func kill(i *io.Input, ac *ActionContext, eventService *EventService) *output {
	mob := ac.getMobBySyntax(mobInRoomSyntax)
	newFight(ac.mob, mob)
	return newOutputToRequestCreator(i, CompletedStatus, "You scream and attack!")
}

func flee(i *io.Input, ac *ActionContext, eventService *EventService) *output {
	ac.mob.fight = nil
	ac.mob.move(ac.mob.room.exits[dice().Intn(len(ac.mob.room.exits))])
	return newOutputToRequestCreator(i, CompletedStatus, "you flee!")
}

func look(i *io.Input, ac *ActionContext, eventService *EventService) *output {
	r := ac.room
	return newOutputToRequestCreator(
		i,
		CompletedStatus,
		fmt.Sprintf(
			"%s\n%s\n%s\n%s%s",
			r.name,
			r.description,
			exitsString(r),
			itemsString(r),
			mobsString(r, ac.mob),
		),
	)
}

func wear(i *io.Input, ac *ActionContext, eventService *EventService) *output {
	for j, item := range ac.mob.items {
		if i.MatchesSubject(item.identifiers) {
			for k, eq := range ac.mob.equipped {
				if eq.position == item.position {
					ac.mob.equipped, ac.mob.items = transferItem(k, ac.mob.equipped, ac.mob.items)
					//i.mob.notify(fmt.Sprintf("You remove %s and put it in your inventory.", eq.String()))
				}
			}
			ac.mob.items, ac.mob.equipped = transferItem(j, ac.mob.items, ac.mob.equipped)
			return newOutputToRequestCreator(i, CompletedStatus, fmt.Sprintf("You wear %s.", item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func remove(i *io.Input, ac *ActionContext, eventService *EventService) *output {
	for j, item := range ac.mob.equipped {
		if i.MatchesSubject(item.identifiers) {
			ac.mob.equipped, ac.mob.items = transferItem(j, ac.mob.equipped, ac.mob.items)
			return newOutputToRequestCreator(i, CompletedStatus, fmt.Sprintf("You remove %s.", item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func get(i *io.Input, ac *ActionContext, eventService *EventService) *output {
	for j, item := range ac.room.items {
		if i.MatchesSubject(item.identifiers) {
			ac.room.items, ac.mob.items = transferItem(j, ac.room.items, ac.mob.items)
			return newOutput(
				i,
				CompletedStatus,
				fmt.Sprintf("You pick up %s", item.String()),
				fmt.Sprintf("%s picks up %s", ac.mob.name, item.String()),
				fmt.Sprintf("%s picks up %s", ac.mob.name, item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func drop(i *io.Input, ac *ActionContext, eventService *EventService) *output {
	for j, item := range ac.mob.items {
		if i.MatchesSubject(item.identifiers) {
			ac.mob.items, ac.room.items = transferItem(j, ac.mob.items, ac.room.items)
			return newOutput(
				i,
				CompletedStatus,
				fmt.Sprintf("You drop %s", item.String()),
				fmt.Sprintf("%s drops %s", ac.mob.name, item.String()),
				fmt.Sprintf("%s drops %s", ac.mob.name, item.String()))
		}
	}
	return newOutputToRequestCreator(
		i,
		FailedStatus,
		"You can't find that.")
}

func move(d direction, i *io.Input, ac *ActionContext, eventService *EventService) *output {
	log.Printf("move: direction: %s, Command: %s", d, i.GetCommand())
	for _, e := range ac.room.exits {
		if e.direction == d {
			eventService.Publish(&Event{
				eventType:MobMoveEventType,
				mob: ac.mob,
				room: e.room,
			})
			return newOutputToRequestCreator(i, CompletedStatus, fmt.Sprintf("You move %s.", d))
		}
	}
	return newOutputToRequestCreator(i, FailedStatus, "Alas, you cannot go that way.")
}
