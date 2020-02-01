package util

import "log"

func Pad(value string, amount int) string {
	log.Print(value)
	for i := len(value); i < amount; i++ {
		value += " "
		log.Print("add", value)
	}
	return value
}
