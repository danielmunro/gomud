package gomud

import (
	"errors"
	"github.com/danielmunro/gomud/io"
	"github.com/danielmunro/gomud/model"
)

type mobRoom struct {
	mob  *model.Mob
	room *model.Room
}

type LocationService struct {
	mobRooms []*mobRoom
}

func newLocationService() *LocationService {
	return &LocationService{
		mobRooms: []*mobRoom{},
	}
}

func (ls *LocationService) changeMobRoom(mob *model.Mob, room *model.Room) {
	for _, mr := range ls.mobRooms {
		if mr.mob == mob {
			mr.room = room
			return
		}
	}
}

func (ls *LocationService) getRoomForMob(mob *model.Mob) *model.Room {
	for _, mr := range ls.mobRooms {
		if mr.mob == mob {
			return mr.room
		}
	}
	return nil
}

func (ls *LocationService) findMobInRoom(buffer *io.Buffer, room *model.Room) (*model.Mob, error) {
	for _, mr := range ls.mobRooms {
		if buffer.MatchesSubject(mr.mob.Identifiers) && mr.room == room {
			return mr.mob, nil
		}
	}
	return nil, errors.New("no mob found")
}

func (ls *LocationService) getMobsInRoom(room *model.Room) []*model.Mob {
	var mobs []*model.Mob
	for _, mr := range ls.mobRooms {
		if mr.room == room {
			mobs = append(mobs, mr.mob)
		}
	}
	return mobs
}

func (ls *LocationService) countMobsInRoom(mob *model.Mob, room *model.Room) int {
	amount := 0
	for _, mr := range ls.mobRooms {
		if mr.mob.ID == mob.ID && mr.room == room {
			amount++
		}
	}
	return amount
}

func (ls *LocationService) countMobsInGame(mob *model.Mob) int {
	amount := 0
	for _, mr := range ls.mobRooms {
		if mr.mob.ID == mob.ID {
			amount++
		}
	}
	return amount
}

func (ls *LocationService) spawnMobToRoom(mob *model.Mob, room *model.Room) {
	ls.mobRooms = append(ls.mobRooms, &mobRoom{
		mob,
		room,
	})
}
