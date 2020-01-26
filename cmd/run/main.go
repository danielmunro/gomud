package main

import (
	"github.com/danielmunro/gomud"
	"github.com/danielmunro/gomud/io"
)

func main() {
	svc := gomud.NewGameService(io.NewServer(8080))
	svc.CreateFixtures()
	svc.StartServer()
}
