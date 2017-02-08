package gomud

type job string

const (
	mage        job = "mage"
	warrior     job = "warrior"
	thief       job = "thief"
	cleric      job = "cleric"
	uninitiated job = "uninitiated"
)

func jobAttributes(j job) *attributes {
	switch j {
	case mage:
		return newAttributes(map[attribute]int{aInt: 2, aWis: 1, aDex: 1})
	case warrior:
		return newAttributes(map[attribute]int{aStr: 2, aCon: 1, aDex: 1})
	case cleric:
		return newAttributes(map[attribute]int{aWis: 2, aInt: 1, aCon: 1})
	case thief:
		return newAttributes(map[attribute]int{aDex: 2, aCon: 1, aStr: 1})
	default:
		return &attributes{}
	}
}
