package gomud

type RoomService struct {
	rooms []*room
}

func newRoomService() *RoomService {
	return &RoomService{
		rooms: []*room{},
	}
}

func (rs *RoomService) addRoom(room *room) {
	rs.rooms = append(rs.rooms, room)
}
