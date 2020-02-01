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
	return ac.CreateOutput(message.GetKillMessage(ac.mob, target))
}

func flee(ac *ActionContext, actionService *ActionService) *io.Output {
	actionService.Publish(NewEvent(FleeEventType, ac.mob, ac.room))
	return ac.CreateOutput(message.GetFleeMessage(ac.mob))
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
		return ac.CreateOutput(message.GetRemoveAndWearMessage(ac.mob, item, equipped))
	}
	return ac.CreateOutput(message.GetWearMessage(ac.mob, item))
}

func remove(ac *ActionContext, _ *ActionService) *io.Output {
	item := ac.getItemBySyntax(itemEquippedSyntax)
	ac.mob.Equipped, ac.mob.Items = transferItem(item, ac.mob.Equipped, ac.mob.Items)
	return ac.CreateOutput(message.GetRemoveMessage(ac.mob, item))
}

func get(ac *ActionContext, _ *ActionService) *io.Output {
	item := ac.getItemBySyntax(itemInRoomSyntax)
	ac.room.Items, ac.mob.Items = transferItem(item, ac.room.Items, ac.mob.Items)
	return ac.CreateOutput(message.GetGetMessage(ac.mob, item))
}

func drop(ac *ActionContext, _ *ActionService) *io.Output {
	item := ac.getItemBySyntax(itemInInventorySyntax)
	ac.mob.Items, ac.room.Items = transferItem(item, ac.mob.Items, ac.room.Items)
	return ac.CreateOutput(message.GetDropMessage(ac.mob, item))
}

func move(d model.Direction, ac *ActionContext, actionService *ActionService) *io.Output {
	log.Printf("move: direction: %s, Command: %s", d, ac.buffer.GetCommand())
	exit := ac.getExitBySyntax(exitDirectionSyntax)
	actionService.Publish(&Event{
		eventType: MobMoveEventType,
		mob:       ac.mob,
		room:      exit.Room,
	})
	return ac.CreateOutput(message.GetMoveMessage(ac.mob, d))
}

func inventory(ac *ActionContext, _ *ActionService) *io.Output {
	buf := "your inventory:\n"
	for _, i := range ac.mob.Items {
		buf += i.String() + "\n"
	}
	return ac.buffer.CreateOutputToRequestCreator(buf)
}

func sit(ac *ActionContext, _ *ActionService) *io.Output {
	wasSleeping := ac.mob.IsSleeping()
	ac.mob.SetSittingDisposition()
	return ac.CreateOutput(message.GetSitMessage(ac.mob, wasSleeping))
}

func sleep(ac *ActionContext, _ *ActionService) *io.Output {
	ac.mob.SetSleepingDisposition()
	return ac.CreateOutput(message.GetSleepMessage(ac.mob))
}

func wake(ac *ActionContext, _ *ActionService) *io.Output {
	wasSleeping := ac.mob.IsSleeping()
	ac.mob.SetStandingDisposition()
	return ac.CreateOutput(message.GetWakeMessage(ac.mob, wasSleeping))
}

func list(ac *ActionContext, _ *ActionService) *io.Output {
	merchant := ac.getMobBySyntax(merchantInRoomSyntax)
	buffer := fmt.Sprintf("%s is selling:\n", merchant.String())
	for _, i := range merchant.Items {
		buffer += fmt.Sprintf("[%d %d] %s\n", i.Level, i.Value, i.String())
	}
	return ac.buffer.CreateOutputToRequestCreator(buffer)
}

func buy(ac *ActionContext, _ *ActionService) *io.Output {
	merchant := ac.getMobBySyntax(merchantInRoomSyntax)
	item := ac.getItemBySyntax(itemInTargetInventorySyntax)
	if ac.mob.Gold < item.Value {
		return ac.buffer.CreateOutputToRequestCreator("You cannot afford that.")
	}
	ac.mob.Gold -= item.Value
	merchant.Gold += item.Value
	if item.IsStoreItem {
		itemToGive := item
		ac.mob.Items = append(ac.mob.Items, itemToGive)
	} else {
		transferItem(item, merchant.Items, ac.mob.Items)
	}
	return ac.CreateOutput(message.GetBuyMessage(ac.mob, merchant, item))
}

func sell(ac *ActionContext, _ *ActionService) *io.Output {
	merchant := ac.getMobBySyntax(merchantInRoomSyntax)
	item := ac.getItemBySyntax(itemInTargetInventorySyntax)
	if merchant.Gold < item.Value {
		return ac.buffer.CreateOutputToRequestCreator("They cannot afford that.")
	}
	if item.IsStoreItem {
		return ac.buffer.CreateOutputToRequestCreator("They are not interested in that.")
	}
	ac.mob.Gold += item.Value
	merchant.Gold -= item.Value
	transferItem(item, ac.mob.Items, merchant.Items)
	return ac.CreateOutput(message.GetBuyMessage(ac.mob, merchant, item))
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
