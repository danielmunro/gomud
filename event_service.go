package gomud


type EventService struct {
	observers []*Observer
}

func NewEventService(gs *GameService) *EventService {
	return &EventService{
		observers: newObservers(gs),
	}
}

func (es *EventService) Publish(event *Event) {
	for _, observer := range es.observers {
		if observer.eventType == event.eventType {
			observer.call(event)
		}
	}
}
