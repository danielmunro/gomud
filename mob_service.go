package gomud

type MobService struct {
	mobResets []*MobReset
	fights []*Fight
}

func NewMobService() *MobService {
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

func (ms *MobService) AddFight(fight *Fight) {
	ms.fights = append(ms.fights, fight)
}

func (ms *MobService) ProceedFights() {
	for i, f := range ms.fights {
		f.Proceed()
		if f.IsEnded() {
			ms.fights = append(ms.fights[:i], ms.fights[i+1:]...)
		}
	}
}

func (ms *MobService) addMobReset(mobReset *MobReset) {
	ms.mobResets = append(ms.mobResets, mobReset)
}
