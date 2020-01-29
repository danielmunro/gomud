package gomud

type Event struct {
	eventType EventType
	mob *Mob
	room *Room
	target *Mob
}

func NewTargetEvent(eventType EventType, mob *Mob, target *Mob, room *Room) *Event {
	return &Event{
		eventType,
		mob,
		room,
		target,
	}
}

func NewEvent(eventType EventType, mob *Mob, room *Room) *Event {
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
