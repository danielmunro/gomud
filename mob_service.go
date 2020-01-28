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

func (ms *MobService) addMobReset(mobReset *MobReset) {
	ms.mobResets = append(ms.mobResets, mobReset)
}
