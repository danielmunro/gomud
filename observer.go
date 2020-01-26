package gomud

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

func newObservers(gs *GameService) []*Observer {
	return []*Observer{
		newMobMoveObserver(gs.locationService),
	}
}
