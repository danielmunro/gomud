package gomud

import "math/rand"

type EventObserver func(event *Event)

type Observer struct {
	eventType     EventType
	eventObserver EventObserver
}

func NewObserver(eventType EventType, eventObserver EventObserver) *Observer {
	return &Observer{
		eventType,
		eventObserver,
	}
}

func newMobMoveObserver(locationService *LocationService) *Observer {
	return NewObserver(MobMoveEventType, func(event *Event) {
		locationService.changeMobRoom(event.mob, event.room)
	})
}

func newFleeObserver(locationService *LocationService, mobService *MobService) *Observer {
	return NewObserver(FleeEventType, func(event *Event) {
		mobService.EndFightForMob(event.mob)
		newRoom := event.room.exits[rand.Intn(len(event.room.exits))].room
		locationService.changeMobRoom(event.mob, newRoom)
	})
}

func newObservers(ls *LocationService, ms *MobService) []*Observer {
	return []*Observer{
		newMobMoveObserver(ls),
		newFleeObserver(ls, ms),
	}
}
