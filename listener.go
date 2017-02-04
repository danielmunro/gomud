package gomud

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

var mobs []*mob
var clients []*client

// Listen ...
func Listen(port int) error {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	defer ln.Close()
	listener := make(chan *client)
	go inputLoop(listener)
	go timing()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go func(c *client) {
			clients = append(clients, c)
			mobs = append(mobs, c.mob)
			look(&input{client: c})
			for {
				listener <- c.read()
			}
		}(newClient(conn, startRoom()))
	}
}

func inputLoop(listener chan *client) {
	for {
		select {
		case c := <-listener:
			parse(newInput(c, strings.Split(c.message, " ")))
		}
	}
}

func timing() {
	const (
		tickLen int = 15
	)
	var pulse int64
	var tick int
	for {
		p := time.Now().Unix()
		if p > pulse {
			for _, m := range mobs {
				if d20() == 1 {
					roleCheck(m)
				}
			}
			pulse = p
			tick++
			if tick > tickLen {
				log.Println(fmt.Sprintf("tick at %d", p))
				tick = 0
				for _, m := range mobs {
					regen(m)
				}
				for _, c := range clients {
					c.writePrompt("")
				}
			}
		}
	}
}

func roleCheck(m *mob) {
	for _, r := range m.roles {
		if r == mobile {
			m.roam()
		} else if r == scavenger {
			m.scavenge()
		}
	}
}

func regen(m *mob) {

}
