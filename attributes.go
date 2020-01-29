package gomud

import "github.com/jinzhu/gorm"

type attribute string

const (
	aHp       attribute = "hp"
	aMana     attribute = "mana"
	aMv       attribute = "mv"
	aStr      attribute = "str"
	aInt      attribute = "int"
	aWis      attribute = "wis"
	aDex      attribute = "dex"
	aCon      attribute = "con"
	aHit      attribute = "hit"
	aDam      attribute = "dam"
	aAcBash   attribute = "acBash"
	aAcSlash  attribute = "acSlash"
	aAcPierce attribute = "acPierce"
	aAcMagic  attribute = "acMagic"
)

type attributes struct {
	gorm.Model
	values map[attribute]int
}

func newStats(str int, i int, wis int, dex int, con int) *attributes {
	return &attributes{
		values: map[attribute]int{
			aStr: str,
			aInt: i,
			aWis: wis,
			aDex: dex,
			aCon: con,
		},
	}
}

func (a *attributes) value(attr attribute) int {
	i, ok := a.values[attr]

	if ok {
		return i
	}

	return 0
}
