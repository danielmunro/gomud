package main

import "github.com/danielmunro/gomud"

func main() {
	gomud.NewListener().Listen(8080)
}
