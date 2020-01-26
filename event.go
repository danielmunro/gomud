package gomud

type Event struct {
	eventType EventType
	mob *mob
	room *room
}
