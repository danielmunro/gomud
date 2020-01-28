package gomud


type EventService struct {
	observers []*Observer
}

func NewEventService(ls *LocationService, ms *MobService) *EventService {
	return &EventService{
		observers: newObservers(ls, ms),
	}
}

func (es *EventService) Publish(event *Event) {
	for _, observer := range es.observers {
		if observer.eventType == event.eventType {
			observer.eventObserver(event)
		}
	}
}
