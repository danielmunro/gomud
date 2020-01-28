package gomud

import "math/rand"

type Observer struct {
	eventType EventType
	call func(event *Event)
}

func newMobMoveObserver(locationService *LocationService) *Observer {
	return &Observer{
		eventType: MobMoveEventType,
		call: func(event *Event) {
			locationService.changeMobRoom(event.mob, event.room)
		},
	}
}

func newFleeObserver(locationService *LocationService, mobService *MobService) *Observer {
	return &Observer{
		eventType: FleeEventType,
		call: func(event *Event) {
			mobService.EndFightForMob(event.mob)
			newRoom := event.room.exits[rand.Intn(len(event.room.exits))].room
			locationService.changeMobRoom(event.mob, newRoom)
		},
	}
}

func newObservers(gs *GameService) []*Observer {
	return []*Observer{
		newMobMoveObserver(gs.locationService),
		newFleeObserver(gs.locationService, gs.mobService),
	}
}
