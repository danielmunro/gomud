package gomud

type mobRoom struct {
	mob *Mob
	room *room
}

type LocationService struct {
	mobRooms []*mobRoom
}

func newLocationService() *LocationService {
	return &LocationService{
		mobRooms:[]*mobRoom{},
	}
}

func (ls *LocationService) changeMobRoom(mob *Mob, room *room) {
	for _, mr := range ls.mobRooms {
		if mr.mob == mob {
			mr.room = room
			return
		}
	}
}

func (ls *LocationService) getRoomForMob(mob *Mob) *room {
	for _, mr := range ls.mobRooms {
		if mr.mob == mob {
			return mr.room
		}
	}
	return nil
}

func (ls *LocationService) findMobInRoom(i *input) *Mob {
	for _, mr := range ls.mobRooms {
		if i.matchesSubject(mr.mob.identifiers) && i.mob.room == mr.room {
			return mr.mob
		}
	}
	return nil
}

func (ls *LocationService) countMobsInRoom(mob *Mob, room *room) int {
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

func (ls *LocationService) spawnMobToRoom(mob *Mob, room *room) {
	ls.mobRooms = append(ls.mobRooms, &mobRoom{
		mob,
		room,
	})
}
