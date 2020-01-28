package gomud

import (
	"errors"
	"github.com/danielmunro/gomud/io"
)

type mobRoom struct {
	mob *Mob
	room *Room
}

type LocationService struct {
	mobRooms []*mobRoom
}

func newLocationService() *LocationService {
	return &LocationService{
		mobRooms:[]*mobRoom{},
	}
}

func (ls *LocationService) changeMobRoom(mob *Mob, room *Room) {
	for _, mr := range ls.mobRooms {
		if mr.mob == mob {
			mr.room = room
			return
		}
	}
}

func (ls *LocationService) getRoomForMob(mob *Mob) *Room {
	for _, mr := range ls.mobRooms {
		if mr.mob == mob {
			return mr.room
		}
	}
	return nil
}

func (ls *LocationService) findMobInRoom(buffer *io.Buffer, room *Room) (*Mob, error) {
	for _, mr := range ls.mobRooms {
		if buffer.MatchesSubject(mr.mob.identifiers) && mr.room == room {
			return mr.mob, nil
		}
	}
	return nil, errors.New("no mob found")
}

func (ls *LocationService) getMobsInRoom(room *Room) []*Mob {
	var mobs []*Mob
	for _, mr := range ls.mobRooms {
		if mr.room == room {
			mobs = append(mobs, mr.mob)
		}
	}
	return mobs
}

func (ls *LocationService) countMobsInRoom(mob *Mob, room *Room) int {
	amount := 0
	for _, mr := range ls.mobRooms {
		if mr.mob.ID == mob.ID && mr.room == room {
			amount++
		}
	}
	return amount
}

func (ls *LocationService) countMobsInGame(mob *Mob) int {
	amount := 0
	for _, mr := range ls.mobRooms {
		if mr.mob.ID == mob.ID {
			amount++
		}
	}
	return amount
}

func (ls *LocationService) spawnMobToRoom(mob *Mob, room *Room) {
	ls.mobRooms = append(ls.mobRooms, &mobRoom{
		mob,
		room,
	})
}
