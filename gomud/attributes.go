package gomud

type AC struct {
	Pierce, Bash, Slash, Magic float64
}

type Vitals struct {
	Hp, Mana, Mv float64
}

type Stats struct {
	Str, Int, Wis, Dex, Con, Luck int
}

type Attributes struct {
	Hit, Dam float64
	Stats *Stats
	Vitals *Vitals
	AC *AC
}
