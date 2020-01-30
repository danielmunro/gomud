package gomud

import "github.com/danielmunro/gomud/model"

type RoomService struct {
	rooms []*model.Room
}

func newRoomService() *RoomService {
	return &RoomService{
		rooms: []*model.Room{},
	}
}

func (rs *RoomService) addRoom(room *model.Room) {
	rs.rooms = append(rs.rooms, room)
}
