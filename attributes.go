package gomud

import "github.com/jinzhu/gorm"

const startHp = 20
const startMana = 100
const startMv = 100

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

type Attributes struct {
	gorm.Model
	Values map[attribute]int
}

func NewStartingAttrs() *Attributes {
	return &Attributes{
		Values: map[attribute]int{
			aHp: startHp,
			aMana: startMana,
			aMv: startMv,
		},
	}
}

func NewStats(str int, _int int, wis int, dex int, con int) *Attributes {
	return &Attributes{
		Values: map[attribute]int{
			aStr: str,
			aInt: _int,
			aWis: wis,
			aDex: dex,
			aCon: con,
		},
	}
}

func (a *Attributes) Value(attr attribute) int {
	i, ok := a.Values[attr]

	if ok {
		return i
	}

	return 0
}
