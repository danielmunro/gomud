package util

import (
	"math/rand"
	"time"
)

func d20() int {
	return dInt(20)
}

func dInt(d int) int {
	return Dice().Intn(d-1) + 1
}

func Dice() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
