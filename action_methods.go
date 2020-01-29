package gomud

import (
	"fmt"
	"github.com/danielmunro/gomud/io"
	"log"
)

func kill(ac *ActionContext, actionService *ActionService) *io.Output {
	mob := ac.getMobBySyntax(mobInRoomSyntax)
	actionService.Publish(NewTargetEvent(AttackEventType, ac.mob, mob, ac.room))
	return ac.buffer.CreateOutput(
		fmt.Sprintf("You scream and attack %s!", mob.String()),
		fmt.Sprintf("%s screams and attacks you!", ac.mob.String()),
		fmt.Sprintf("%s screams and attacks %s!", ac.mob.String(), mob.String()))
}

func flee(ac *ActionContext, actionService *ActionService) *io.Output {
	actionService.Publish(NewEvent(FleeEventType, ac.mob, ac.room))
	return ac.buffer.CreateOutputToRequestCreator("you flee!")
}

func look(ac *ActionContext, actionService *ActionService) *io.Output {
	r := ac.room
	return ac.buffer.CreateOutputToRequestCreator(fmt.Sprintf(
		"%s\n%s\n%s\n%s%s",
		r.name,
		r.description,
		exitsString(r),
		itemsString(r),
		mobsString(actionService.GetMobsForRoomAndObserver(ac.room, ac.mob)),
	))
}

func wear(ac *ActionContext, _ *ActionService) *io.Output {
	buf := ""
	item := ac.getItemBySyntax(itemInInventorySyntax)
	for k, eq := range ac.mob.equipped {
		if eq.position == item.position {
			ac.mob.equipped, ac.mob.items = transferItemByIndex(k, ac.mob.equipped, ac.mob.items)
			buf = fmt.Sprintf("You remove %s and put it in your inventory. ", eq.String())
		}
	}
	ac.mob.items, ac.mob.equipped = transferItem(item, ac.mob.items, ac.mob.equipped)
	return ac.buffer.CreateOutputToRequestCreator(buf + fmt.Sprintf("You wear %s.", item.String()))
}

func remove(ac *ActionContext, _ *ActionService) *io.Output {
	item := ac.getItemBySyntax(itemEquippedSyntax)
	ac.mob.equipped, ac.mob.items = transferItem(item, ac.mob.equipped, ac.mob.items)
	return ac.buffer.CreateOutputToRequestCreator(fmt.Sprintf("You remove %s.", item.String()))
}

func get(ac *ActionContext, _ *ActionService) *io.Output {
	item := ac.getItemBySyntax(itemInRoomSyntax)
	ac.room.items, ac.mob.items = transferItem(item, ac.room.items, ac.mob.items)
	return ac.buffer.CreateOutput(
		fmt.Sprintf("You pick up %s", item.String()),
		fmt.Sprintf("%s picks up %s", ac.mob.name, item.String()),
		fmt.Sprintf("%s picks up %s", ac.mob.name, item.String()))
}

func drop(ac *ActionContext, _ *ActionService) *io.Output {
	item := ac.getItemBySyntax(itemInInventorySyntax)
	ac.mob.items, ac.room.items = transferItem(item, ac.mob.items, ac.room.items)
	return ac.buffer.CreateOutput(
		fmt.Sprintf("You drop %s", item.String()),
		fmt.Sprintf("%s drops %s", ac.mob.name, item.String()),
		fmt.Sprintf("%s drops %s", ac.mob.name, item.String()))
}

func move(d direction, ac *ActionContext, actionService *ActionService) *io.Output {
	log.Printf("move: direction: %s, Command: %s", d, ac.buffer.GetCommand())
	exit := ac.getExitBySyntax(exitDirectionSyntax)
	actionService.Publish(&Event{
		eventType:MobMoveEventType,
		mob: ac.mob,
		room: exit.room,
	})
	return ac.buffer.CreateOutputToRequestCreator(fmt.Sprintf("You move %s.", d))
}

func inventory(ac *ActionContext, _ *ActionService) *io.Output {
	buf := "your inventory:\n"
	for _, i := range ac.mob.items {
		buf += i.String() + "\n"
	}
	return ac.buffer.CreateOutputToRequestCreator(buf)
}

func transferItem(item *item, from []*item, to []*item) ([]*item, []*item) {
	for i, x := range from {
		if x == item {
			from = append(from[0:i], from[i+1:]...)
			to = append(to, item)
		}
	}

	return from, to
}

func transferItemByIndex(i int, from []*item, to []*item) ([]*item, []*item) {
	item := from[i]
	from = append(from[0:i], from[i+1:]...)
	to = append(to, item)

	return from, to
}

func exitsString(r *Room) string {
	var exits string

	for _, e := range r.exits {
		exits = fmt.Sprintf("%s%s", exits, string(e.direction[0]))
	}

	return fmt.Sprintf("[%s]", exits)
}

func mobsString(mobs []*Mob) string {
	var buf string

	for _, m := range mobs {
		buf = fmt.Sprintf("%s is here.\n%s", m.String(), buf)
	}

	return buf
}

func itemsString(r *Room) string {
	var items string

	for _, i := range r.items {
		items = fmt.Sprintf("%s is here.\n%s", i.String(), items)
	}

	return items
}
