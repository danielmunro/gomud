package message

import (
	"fmt"
	"github.com/danielmunro/gomud/model"
)

const (
	ErrorDispositionMismatch = "you must be %s to do that."
	ErrorInputNotUnderstood = "what was that?"
)

func GetFleeMessage(mobName string) *Message {
	return &Message{
		ToRequestCreator: "you flee!",
		ToTarget:         fmt.Sprintf("%s flees!", mobName),
		ToObservers:      fmt.Sprintf("%s flees!", mobName),
	}
}

func GetKillMessage(mobName string, targetName string) *Message {
	return &Message{
		ToRequestCreator: fmt.Sprintf("You scream and attack %s!", targetName),
		ToTarget: fmt.Sprintf("%s screams and attacks you!", mobName),
		ToObservers: fmt.Sprintf("%s screams and attacks %s!", mobName, targetName),
	}
}

func GetRemoveAndWearMessage(mobName string, toEquipName string, removeName string) *Message {
	return &Message{
		fmt.Sprintf("You remove %s and put it in your inventory. You wear %s.", removeName, toEquipName),
		fmt.Sprintf("%s removes %s and puts it in their inventory. They wear %s.", mobName, removeName, toEquipName),
		fmt.Sprintf("%s removes %s and puts it in their inventory. They wear %s.", mobName, removeName, toEquipName),
	}
}

func GetRemoveMessage(mobName string, removeName string) *Message {
	return &Message{
		fmt.Sprintf("You take %s off and put it in your inventory.", removeName),
		fmt.Sprintf("%s takes %s off and puts it in their inventory.", mobName, removeName),
		fmt.Sprintf("%s takes %s and puts it in their inventory.", mobName, removeName),
	}
}

func GetGetMessage(mob *model.Mob, item *model.Item) *Message {
	name := mob.Name
	pronoun := mob.GetGenderPronoun()
	itemName := item.String()
	return &Message{
		fmt.Sprintf("You pick up %s and put it in your inventory.", itemName),
		fmt.Sprintf("%s picks up %s and puts it in %s inventory.", name, itemName, pronoun),
		fmt.Sprintf("%s picks up %s and puts it in %s inventory.", name, itemName, pronoun),
	}
}

func GetWearMessage(mobName string, toEquipName string) *Message {
	return &Message{
		fmt.Sprintf("You wear %s.", toEquipName),
		fmt.Sprintf("%s wears %s.", mobName, toEquipName),
		fmt.Sprintf("%s wears %s.", mobName, toEquipName),
	}
}
