package gomud

type EventType string

const (
	MobMoveEventType EventType = "mob move"
	FleeEventType    EventType = "flee"
	LookEventType    EventType = "look"
	AttackEventType  EventType = "attack"
	PulseEventType   EventType = "pulse"
)
