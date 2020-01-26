package gomud

type MobReset struct {
	mob *Mob
	room *Room
	maxInRoom int
	maxInGame int
}

func NewMobReset(mob *Mob, room *Room, maxInRoom int, maxInGame int) *MobReset {
	return &MobReset{
		mob,
		room,
		maxInRoom,
		maxInGame,
	}
}
