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
	eventService.Publish(NewEvent(FleeEventType, ac.mob, ac.room))
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
	buf := ""
	item := ac.getItemBySyntax(itemInInventorySyntax)
	for k, eq := range ac.mob.equipped {
		if eq.position == item.position {
			ac.mob.equipped, ac.mob.items = transferItemByIndex(k, ac.mob.equipped, ac.mob.items)
			buf = fmt.Sprintf("You remove %s and put it in your inventory. ", eq.String())
		}
	}
	ac.mob.items, ac.mob.equipped = transferItem(item, ac.mob.items, ac.mob.equipped)
	return io.NewOutputToRequestCreator(b, io.CompletedStatus, buf + fmt.Sprintf("You wear %s.", item.String()))
}

func remove(b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	item := ac.getItemBySyntax(itemEquippedSyntax)
	ac.mob.equipped, ac.mob.items = transferItem(item, ac.mob.equipped, ac.mob.items)
	return io.NewOutputToRequestCreator(b, io.CompletedStatus, fmt.Sprintf("You remove %s.", item.String()))
}

func get(b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	item := ac.getItemBySyntax(itemInRoomSyntax)
	ac.room.items, ac.mob.items = transferItem(item, ac.room.items, ac.mob.items)
	return io.NewOutput(
		b,
		io.CompletedStatus,
		fmt.Sprintf("You pick up %s", item.String()),
		fmt.Sprintf("%s picks up %s", ac.mob.name, item.String()),
		fmt.Sprintf("%s picks up %s", ac.mob.name, item.String()))
}

func drop(b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	item := ac.getItemBySyntax(itemInInventorySyntax)
	ac.mob.items, ac.room.items = transferItem(item, ac.mob.items, ac.room.items)
	return io.NewOutput(
		b,
		io.CompletedStatus,
		fmt.Sprintf("You drop %s", item.String()),
		fmt.Sprintf("%s drops %s", ac.mob.name, item.String()),
		fmt.Sprintf("%s drops %s", ac.mob.name, item.String()))
}

func move(d direction, b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	log.Printf("move: direction: %s, Command: %s", d, b.GetCommand())
	exit := ac.getExitBySyntax(exitDirectionSyntax)
	eventService.Publish(&Event{
		eventType:MobMoveEventType,
		mob: ac.mob,
		room: exit.room,
	})
	return io.NewOutputToRequestCreator(b, io.CompletedStatus, fmt.Sprintf("You move %s.", d))
}

func inventory(b *io.Buffer, ac *ActionContext, eventService *EventService) *io.Output {
	buf := "your inventory:\n"
	for _, i := range ac.mob.items {
		buf += i.String() + "\n"
	}
	return io.NewOutputToRequestCreator(b, io.CompletedStatus, buf)
}
