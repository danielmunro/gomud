package gomud

type Event struct {
	eventType EventType
	mob *Mob
	room *Room
}

func NewEvent(eventType EventType, mob *Mob, room *Room) *Event {
	return &Event{
		eventType,
		mob,
		room,
	}
}
