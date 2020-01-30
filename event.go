package gomud

import "github.com/danielmunro/gomud/model"

type Event struct {
	eventType EventType
	mob *model.Mob
	room *model.Room
	target *model.Mob
}

func NewTargetEvent(eventType EventType, mob *model.Mob, target *model.Mob, room *model.Room) *Event {
	return &Event{
		eventType,
		mob,
		room,
		target,
	}
}

func NewEvent(eventType EventType, mob *model.Mob, room *model.Room) *Event {
	return &Event{
		eventType,
		mob,
		room,
		nil,
	}
}

func NewSystemEvent(eventType EventType) *Event {
	return &Event {
		eventType,
		nil,
		nil,
		nil,
	}
}
