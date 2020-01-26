package gomud

type Event struct {
	eventType EventType
	mob *Mob
	room *room
}
