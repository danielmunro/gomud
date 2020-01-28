package main

import (
	"github.com/danielmunro/gomud"
	"github.com/danielmunro/gomud/io"
	"log"
)

func main() {
	server, err := io.NewServer(8080)
	if err != nil {
		log.Fatal(err)
	}
	svc := gomud.NewGameService(server)
	svc.CreateFixtures()
	svc.StartServer()
}
