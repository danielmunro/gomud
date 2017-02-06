package gomud

import "github.com/jinzhu/gorm"

type attributes struct {
	gorm.Model
	hp       int
	mana     int
	mv       int
	str      int
	int      int
	wis      int
	dex      int
	con      int
	hit      int
	dam      int
	acBash   int
	acSlash  int
	acPierce int
	acMagic  int
}
