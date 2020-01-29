package gomud

type MobService struct {
	mobResets []*MobReset
	fights []*fight
}

func newMobService() *MobService {
	return &MobService{}
}

func (ms *MobService) EndFightForMob(mob *Mob) {
	for _, f := range ms.fights {
		if f.IncludesMob(mob) {
			f.End()
			return
		}
	}
}

func (ms *MobService) AddFight(fight *fight) {
	ms.fights = append(ms.fights, fight)
}

func (ms *MobService) addMobReset(mobReset *MobReset) {
	ms.mobResets = append(ms.mobResets, mobReset)
}
