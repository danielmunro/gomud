package gomud

import "github.com/danielmunro/gomud/model"

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

func (as *ActionService) GetMobsForRoomAndObserver(room *model.Room, mob *model.Mob) []*model.Mob {
	var mobs []*model.Mob
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
