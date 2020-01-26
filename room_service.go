package gomud

type RoomService struct {
	rooms []*Room
}

func newRoomService() *RoomService {
	return &RoomService{
		rooms: []*Room{},
	}
}

func (rs *RoomService) addRoom(room *Room) {
	rs.rooms = append(rs.rooms, room)
}
