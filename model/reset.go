package model

type MobReset struct {
	Mob       *Mob
	Room      *Room
	MaxInRoom int
	MaxInGame int
}

func NewMobReset(mob *Mob, room *Room, maxInRoom int, maxInGame int) *MobReset {
	return &MobReset{
		mob,
		room,
		maxInRoom,
		maxInGame,
	}
}
