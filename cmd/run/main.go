package main

import (
	"github.com/danielmunro/gomud"
)

func main() {
	svc := gomud.NewGameService(gomud.NewServer(8080))
	svc.CreateFixtures()
	svc.StartServer()
}
