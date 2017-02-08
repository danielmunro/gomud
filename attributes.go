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

func newAttributes(v map[attribute]int) *attributes {
	return &attributes{
		values: v,
	}
}

func (a *attributes) a(attr attribute) int {
	i, ok := a.values[attr]

	if ok {
		return i
	}

	return 0
}
