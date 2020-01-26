package gomud

import (
	"fmt"
	"github.com/danielmunro/gomud/io"
	"log"
)

func kill(b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	mob := ac.getMobBySyntax(mobInRoomSyntax)
	newFight(ac.mob, mob)
	return io.NewOutputToRequestCreator(b, io.CompletedStatus, "You scream and attack!")
}

func flee(b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	ac.mob.fight = nil
	ac.mob.move(ac.mob.room.exits[dice().Intn(len(ac.mob.room.exits))])
	return io.NewOutputToRequestCreator(b, io.CompletedStatus, "you flee!")
}

func look(b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	r := ac.room
	return io.NewOutputToRequestCreator(
		b,
		io.CompletedStatus,
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

func wear(b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	for j, item := range ac.mob.items {
		if b.MatchesSubject(item.identifiers) {
			for k, eq := range ac.mob.equipped {
				if eq.position == item.position {
					ac.mob.equipped, ac.mob.items = transferItem(k, ac.mob.equipped, ac.mob.items)
					//i.mob.notify(fmt.Sprintf("You remove %s and put it in your inventory.", eq.String()))
				}
			}
			ac.mob.items, ac.mob.equipped = transferItem(j, ac.mob.items, ac.mob.equipped)
			return io.NewOutputToRequestCreator(b, io.CompletedStatus, fmt.Sprintf("You wear %s.", item.String()))
		}
	}
	return io.NewOutputToRequestCreator(
		b,
		io.FailedStatus,
		"You can't find that.")
}

func remove(b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	for j, item := range ac.mob.equipped {
		if b.MatchesSubject(item.identifiers) {
			ac.mob.equipped, ac.mob.items = transferItem(j, ac.mob.equipped, ac.mob.items)
			return io.NewOutputToRequestCreator(b, io.CompletedStatus, fmt.Sprintf("You remove %s.", item.String()))
		}
	}
	return io.NewOutputToRequestCreator(
		b,
		io.FailedStatus,
		"You can't find that.")
}

func get(b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	for j, item := range ac.room.items {
		if b.MatchesSubject(item.identifiers) {
			ac.room.items, ac.mob.items = transferItem(j, ac.room.items, ac.mob.items)
			return io.NewOutput(
				b,
				io.CompletedStatus,
				fmt.Sprintf("You pick up %s", item.String()),
				fmt.Sprintf("%s picks up %s", ac.mob.name, item.String()),
				fmt.Sprintf("%s picks up %s", ac.mob.name, item.String()))
		}
	}
	return io.NewOutputToRequestCreator(
		b,
		io.FailedStatus,
		"You can't find that.")
}

func drop(b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	for j, item := range ac.mob.items {
		if b.MatchesSubject(item.identifiers) {
			ac.mob.items, ac.room.items = transferItem(j, ac.mob.items, ac.room.items)
			return io.NewOutput(
				b,
				io.CompletedStatus,
				fmt.Sprintf("You drop %s", item.String()),
				fmt.Sprintf("%s drops %s", ac.mob.name, item.String()),
				fmt.Sprintf("%s drops %s", ac.mob.name, item.String()))
		}
	}
	return io.NewOutputToRequestCreator(
		b,
		io.FailedStatus,
		"You can't find that.")
}

func move(d direction, b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	log.Printf("move: direction: %s, Command: %s", d, b.GetCommand())
	for _, e := range ac.room.exits {
		if e.direction == d {
			eventService.Publish(&Event{
				eventType:MobMoveEventType,
				mob: ac.mob,
				room: e.room,
			})
			return io.NewOutputToRequestCreator(b, io.CompletedStatus, fmt.Sprintf("You move %s.", d))
		}
	}
	return io.NewOutputToRequestCreator(b, io.FailedStatus, "Alas, you cannot go that way.")
}
