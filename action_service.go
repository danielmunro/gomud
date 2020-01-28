package gomud

type ActionService struct {
	locationService *LocationService
	eventService *EventService
}

func newActionService(locationService *LocationService, eventService *EventService) *ActionService {
	return &ActionService{
		locationService,
		eventService,
	}
}

func (as *ActionService) GetMobsForRoomAndObserver(room *Room, mob *Mob) []*Mob {
	var mobs []*Mob
	for _, m := range as.locationService.getMobsInRoom(room) {
		if m != mob {
			mobs = append(mobs, m)
		}
	}
	return mobs
}

func (as *ActionService) Publish(event *Event) {
	as.eventService.Publish(event)
}
