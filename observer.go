package gomud

import (
	"math/rand"
)

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
		newRoom := event.room.Exits[rand.Intn(len(event.room.Exits))].Room
		locationService.changeMobRoom(event.mob, newRoom)
	})
}

func newFightObserver(mobService *MobService) *Observer {
	return NewObserver(AttackEventType, func(event *Event) {
		mobService.AddFight(NewFight(event.mob, event.target))
	})
}

func newProceedFightsObserver(mobService *MobService) *Observer {
	return NewObserver(PulseEventType, func(event *Event) {
		mobService.ProceedFights()
	})
}

func newObservers(ls *LocationService, ms *MobService) []*Observer {
	return []*Observer{
		newMobMoveObserver(ls),
		newFleeObserver(ls, ms),
		newFightObserver(ms),
		newProceedFightsObserver(ms),
	}
}
