package model

import "github.com/jinzhu/gorm"

const startHp = 20
const startMana = 100
const startMv = 100

type Attribute string

const (
	HpAttr     Attribute = "Hp"
	ManaAttr   Attribute = "mana"
	MvAttr     Attribute = "mv"
	StrAttr    Attribute = "str"
	IntAttr    Attribute = "int"
	WisAttr    Attribute = "wis"
	DexAttr    Attribute = "dex"
	ConAttr    Attribute = "con"
	HitAttr    Attribute = "hit"
	DamAttr    Attribute = "dam"
	BashAttr   Attribute = "acBash"
	SlashAttr  Attribute = "acSlash"
	PierceAttr Attribute = "acPierce"
	MagicAttr  Attribute = "acMagic"
)

type Attributes struct {
	gorm.Model
	Values map[Attribute]int
}

func NewStartingAttrs() *Attributes {
	return &Attributes{
		Values: map[Attribute]int{
			HpAttr:   startHp,
			ManaAttr: startMana,
			MvAttr:   startMv,
		},
	}
}

func NewStats(str int, _int int, wis int, dex int, con int) *Attributes {
	return &Attributes{
		Values: map[Attribute]int{
			StrAttr: str,
			IntAttr: _int,
			WisAttr: wis,
			DexAttr: dex,
			ConAttr: con,
		},
	}
}

func (a *Attributes) Value(attr Attribute) int {
	i, ok := a.Values[attr]

	if ok {
		return i
	}

	return 0
}
