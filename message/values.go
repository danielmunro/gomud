package message

import (
	"fmt"
	"github.com/danielmunro/gomud/model"
)

const (
	ErrorDispositionMismatch = "you must be %s to do that."
	ErrorInputNotUnderstood = "what was that?"
)

func GetFleeMessage(mob *model.Mob) *Message {
	mobName := mob.String()
	return &Message{
		ToRequestCreator: "you flee!",
		ToTarget:         fmt.Sprintf("%s flees!", mobName),
		ToObservers:      fmt.Sprintf("%s flees!", mobName),
	}
}

func GetKillMessage(mob *model.Mob, target *model.Mob) *Message {
	mobName := mob.String()
	targetName := target.String()
	return &Message{
		ToRequestCreator: fmt.Sprintf("You scream and attack %s!", targetName),
		ToTarget: fmt.Sprintf("%s screams and attacks you!", mobName),
		ToObservers: fmt.Sprintf("%s screams and attacks %s!", mobName, targetName),
	}
}

func GetRemoveAndWearMessage(mob *model.Mob, toEquip *model.Item, remove *model.Item) *Message {
	mobName := mob.String()
	toEquipName := toEquip.String()
	removeName := remove.String()
	return &Message{
		fmt.Sprintf("You remove %s and put it in your inventory. You wear %s.", removeName, toEquipName),
		fmt.Sprintf("%s removes %s and puts it in their inventory. They wear %s.", mobName, removeName, toEquipName),
		fmt.Sprintf("%s removes %s and puts it in their inventory. They wear %s.", mobName, removeName, toEquipName),
	}
}

func GetRemoveMessage(mob *model.Mob, toRemove *model.Item) *Message {
	mobName := mob.String()
	removeName := toRemove.String()
	return &Message{
		fmt.Sprintf("You take %s off and put it in your inventory.", removeName),
		fmt.Sprintf("%s takes %s off and puts it in their inventory.", mobName, removeName),
		fmt.Sprintf("%s takes %s and puts it in their inventory.", mobName, removeName),
	}
}

func GetGetMessage(mob *model.Mob, item *model.Item) *Message {
	name := mob.String()
	pronoun := mob.GetGenderPronoun()
	itemName := item.String()
	return &Message{
		fmt.Sprintf("You pick up %s and put it in your inventory.", itemName),
		fmt.Sprintf("%s picks up %s and puts it in %s inventory.", name, itemName, pronoun),
		fmt.Sprintf("%s picks up %s and puts it in %s inventory.", name, itemName, pronoun),
	}
}

func GetDropMessage(mob *model.Mob, item *model.Item) *Message {
	mobName := mob.String()
	itemName := item.String()
	return &Message{
		fmt.Sprintf("You drop %s.", itemName),
		fmt.Sprintf("%s drops %s.", mobName, itemName),
		fmt.Sprintf("%s drops %s.", mobName, itemName),
	}
}

func GetWearMessage(mob *model.Mob, toEquip *model.Item) *Message {
	mobName := mob.String()
	toEquipName := toEquip.String()
	return &Message{
		fmt.Sprintf("You wear %s.", toEquipName),
		fmt.Sprintf("%s wears %s.", mobName, toEquipName),
		fmt.Sprintf("%s wears %s.", mobName, toEquipName),
	}
}

func GetSleepMessage(mob *model.Mob) *Message {
	mobName := mob.String()
	return &Message{
		"you lay down and go to sleep.",
		fmt.Sprintf("%s lays down and goes to sleep.", mobName),
		fmt.Sprintf("%s lays down and goes to sleep.", mobName),
	}
}

func GetMoveMessage(mob *model.Mob, direction model.Direction) *Message {
	mobName := mob.String()
	return &Message{
		fmt.Sprintf("You move %s.", direction),
		fmt.Sprintf("%s leaves %s.", mobName, direction),
		fmt.Sprintf("%s leaves %s.", mobName, direction),
	}
}

func GetSitMessage(mob *model.Mob, wasSleeping bool) *Message {
	buf1 := "you "
	buf2 := mob.String() + " "
	if wasSleeping {
		buf1 = "wake up and "
		buf2 = "wakes up and "
	}
	return &Message{
		buf1+"sit down.",
		buf2+"sits down.",
		buf2+"sits down.",
	}
}

func GetWakeMessage(mob *model.Mob, wasSleeping bool) *Message {
	buf1 := "you "
	buf2 := mob.String()
	if wasSleeping {
		buf1 = "wake and "
		buf2 = "wakes and "
	}
	return &Message{
		buf1 + "stand up.",
		buf2 + "stands up.",
		buf2 + "stands up.",
	}
}
