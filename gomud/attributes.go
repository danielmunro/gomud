package gomud

type AC struct {
	Pierce, Bash, Slash, Magic float64
}

type Vitals struct {
	Hp, Mana, Mv float64
}

type Attributes struct {
	Str, Int, Wis, Dex, Con, Luck int
	Hit, Dam float64
	Vitals *Vitals
	AC *AC
}
