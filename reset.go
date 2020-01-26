package gomud

type MobReset struct {
	mob *mob
	room *room
	maxInRoom int
	maxInGame int
}

func NewMobReset(mob *mob, room *room, maxInRoom int, maxInGame int) *MobReset {
	return &MobReset{
		mob,
		room,
		maxInRoom,
		maxInGame,
	}
}
