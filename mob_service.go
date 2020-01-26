package gomud

type MobService struct {
	mobResets []*MobReset
}

func newMobService() *MobService {
	return &MobService{
		mobResets: []*MobReset{},
	}
}

func (ms *MobService) addMobReset(mobReset *MobReset) {
	ms.mobResets = append(ms.mobResets, mobReset)
}
