package gomud

import (
	"fmt"
	"github.com/danielmunro/gomud/io"
	"github.com/danielmunro/gomud/message"
	"github.com/danielmunro/gomud/model"
	"log"
)

func kill(ac *ActionContext, actionService *ActionService) *io.Output {
	target := ac.getMobBySyntax(mobInRoomSyntax)
	actionService.Publish(NewTargetEvent(AttackEventType, ac.mob, target, ac.room))
	return ac.CreateOutputFromMessage(message.GetKillMessage(ac.mob, target))
}

func flee(ac *ActionContext, actionService *ActionService) *io.Output {
	actionService.Publish(NewEvent(FleeEventType, ac.mob, ac.room))
	return ac.CreateOutputFromMessage(message.GetFleeMessage(ac.mob))
}

func look(ac *ActionContext, actionService *ActionService) *io.Output {
	r := ac.room
	return ac.buffer.CreateOutputToRequestCreator(fmt.Sprintf(
		"%s\n%s\n%s\n%s%s",
		r.Name,
		r.Description,
		exitsString(r),
		itemsString(r),
		mobsString(actionService.GetMobsForRoomAndObserver(ac.room, ac.mob)),
	))
}

func wear(ac *ActionContext, _ *ActionService) *io.Output {
	var equipped *model.Item
	item := ac.getItemBySyntax(itemInInventorySyntax)
	for k, eq := range ac.mob.Equipped {
		if eq.Position == item.Position {
			ac.mob.Equipped, ac.mob.Items = transferItemByIndex(k, ac.mob.Equipped, ac.mob.Items)
			equipped = eq
		}
	}
	ac.mob.Items, ac.mob.Equipped = transferItem(item, ac.mob.Items, ac.mob.Equipped)
	if equipped != nil {
		return ac.CreateOutputFromMessage(message.GetRemoveAndWearMessage(ac.mob, item, equipped))
	}
	return ac.CreateOutputFromMessage(message.GetWearMessage(ac.mob, item))
}

func remove(ac *ActionContext, _ *ActionService) *io.Output {
	item := ac.getItemBySyntax(itemEquippedSyntax)
	ac.mob.Equipped, ac.mob.Items = transferItem(item, ac.mob.Equipped, ac.mob.Items)
	return ac.CreateOutputFromMessage(message.GetRemoveMessage(ac.mob, item))
}

func get(ac *ActionContext, _ *ActionService) *io.Output {
	item := ac.getItemBySyntax(itemInRoomSyntax)
	ac.room.Items, ac.mob.Items = transferItem(item, ac.room.Items, ac.mob.Items)
	return ac.CreateOutputFromMessage(message.GetGetMessage(ac.mob, item))
}

func drop(ac *ActionContext, _ *ActionService) *io.Output {
	item := ac.getItemBySyntax(itemInInventorySyntax)
	ac.mob.Items, ac.room.Items = transferItem(item, ac.mob.Items, ac.room.Items)
	return ac.CreateOutputFromMessage(message.GetDropMessage(ac.mob, item))
}

func move(d model.Direction, ac *ActionContext, actionService *ActionService) *io.Output {
	log.Printf("move: direction: %s, Command: %s", d, ac.buffer.GetCommand())
	exit := ac.getExitBySyntax(exitDirectionSyntax)
	actionService.Publish(&Event{
		eventType: MobMoveEventType,
		mob:       ac.mob,
		room:      exit.Room,
	})
	return ac.buffer.CreateOutputToRequestCreator(fmt.Sprintf("You move %s.", d))
}

func inventory(ac *ActionContext, _ *ActionService) *io.Output {
	buf := "your inventory:\n"
	for _, i := range ac.mob.Items {
		buf += i.String() + "\n"
	}
	return ac.buffer.CreateOutputToRequestCreator(buf)
}

func sit(ac *ActionContext, _ *ActionService) *io.Output {
	buf1 := "you "
	buf2 := ac.mob.Name + " "
	if ac.mob.IsSleeping() {
		buf1 = "wake up and "
		buf2 = "wakes up and "
	}
	ac.mob.SetSittingDisposition()
	return ac.buffer.CreateOutput(
		buf1+"sit down.",
		buf2+"sits down.",
		buf2+"sits down.",
	)
}

func sleep(ac *ActionContext, _ *ActionService) *io.Output {
	ac.mob.SetSleepingDisposition()
	return ac.CreateOutputFromMessage(message.GetSleepMessage(ac.mob))
}

func wake(ac *ActionContext, _ *ActionService) *io.Output {
	buf1 := "you "
	buf2 := ac.mob.Name
	if ac.mob.Disposition == model.SleepingDisposition {
		buf1 = "wake and "
		buf2 = "wakes and "
	}
	ac.mob.SetStandingDisposition()
	return ac.buffer.CreateOutput(
		buf1+"stand up.",
		buf2+"stands up.",
		buf2+"stands up.",
	)
}

func transferItem(item *model.Item, from []*model.Item, to []*model.Item) ([]*model.Item, []*model.Item) {
	for i, x := range from {
		if x == item {
			from = append(from[0:i], from[i+1:]...)
			to = append(to, item)
		}
	}

	return from, to
}

func transferItemByIndex(i int, from []*model.Item, to []*model.Item) ([]*model.Item, []*model.Item) {
	item := from[i]
	from = append(from[0:i], from[i+1:]...)
	to = append(to, item)

	return from, to
}

func exitsString(r *model.Room) string {
	var exits string

	for _, e := range r.Exits {
		exits = fmt.Sprintf("%s%s", exits, string(e.Direction[0]))
	}

	return fmt.Sprintf("[%s]", exits)
}

func mobsString(mobs []*model.Mob) string {
	var buf string

	for _, m := range mobs {
		buf = fmt.Sprintf("%s is here.\n%s", m.String(), buf)
	}

	return buf
}

func itemsString(r *model.Room) string {
	var items string

	for _, i := range r.Items {
		items = fmt.Sprintf("%s is here.\n%s", i.String(), items)
	}

	return items
}
