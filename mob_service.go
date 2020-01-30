package gomud

import "github.com/danielmunro/gomud/model"

type MobService struct {
	mobResets []*model.MobReset
	fights []*Fight
}

func NewMobService() *MobService {
	return &MobService{}
}

func (ms *MobService) EndFightForMob(mob *model.Mob) {
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

func (ms *MobService) addMobReset(mobReset *model.MobReset) {
	ms.mobResets = append(ms.mobResets, mobReset)
}
