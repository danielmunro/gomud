package gomud

/*
	AC represents the defenses of a Mob to piercing, bashing,
	slashing, and magical damage.
	Pierce - resistance to piercing damage as a percentage.
	Bash - resistance to bashing damage as a percentage.
	Slash - resistance to slashing damage as a percentage.
	Magic - resistance to magical damage as a percentage.
*/
type AC struct {
	Pierce, Bash, Slash, Magic float64
}

/*
	Vitals respresents the core attributes of a living Mob:
	Hp, Mana, and Movement.
	Hp - the quantity of Hp a Mob has.
	Mana - the quantity of Mana/Magic power a Mob has.
	Mv - the quantity of Movement a Mob has per tick.
*/
type Vitals struct {
	Hp, Mana, Mv float64
}

/*
	Stats defines the core characteristics of a Mob.
	Str - the strength of the Mob.
	Int - the intelligence of the Mob.
	Wis - the wisdom of the Mob.
	Dex - the dexterity of the Mob.
	Con - the constitution of the Mob.
	Luck - the luck of the Mob.
*/
type Stats struct {
	Str, Int, Wis, Dex, Con, Luck int
}

/*
	Attributes encapsulates Vitals, Stats, and AC about a Mob with
	its raw offensive and defensive capabilities.
	Hit - the chance of the Mob hitting it's target.
	Dam - the amount of damage the Mob inflicts on a hit.
	Stats - a Stats containing the core statistics attributes of the Mob.
	Vitals - a Vitals containing the Hp, Mana, and Mv of a Mob.
	AC- an AC containing the defensive values of the Mob.
*/
type Attributes struct {
	Hit, Dam float64
	Stats    *Stats
	Vitals   *Vitals
	AC       *AC
}
