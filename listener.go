package gomud

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type client struct {
	conn    net.Conn
	mob     *mob
	message string
}

var mobs []*mob

func (c *client) read() *client {
	c.message, _ = bufio.NewReader(c.conn).ReadString('\n')
	c.message = strings.Trim(c.message, "\r\n")

	return c
}

func (c *client) write(m string) {
	c.conn.Write([]byte(m))
}

func (c *client) writePrompt(m string) {
	c.write(fmt.Sprintf("%s\n--> ", m))
}

func (c *client) String() string {
	return c.conn.RemoteAddr().String()
}

func newClient(c net.Conn, r *room) *client {
	return &client{
		conn: c,
		mob: &mob{
			name:        "a mob",
			description: "a mob",
			room:        r,
		},
	}
}

// Listen ...
func Listen(port int) error {
	r := scratchWorld()
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	listener := make(chan *client)
	clients := make([]*client, 0)

	go func() {
		for {
			select {
			case c := <-listener:
				parse(newInput(c, strings.Split(c.message, " ")))
			}
		}
	}()

	go func() {
		var pulse int64
		var tick int
		for {
			p := time.Now().Unix()
			if p > pulse {
				for _, m := range mobs {
					if d20() == 1 {
						roleCheck(m)
					}
					//regen(m)
				}
				pulse = p
				tick++
				if tick > 15 {
					log.Println(fmt.Sprintf("tick at %d", p))
					tick = 0
				}
			}
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go func(c *client) {
			c.mob.client = c
			clients = append(clients, c)
			mobs = append(mobs, c.mob)
			look(&input{client: c})
			for {
				listener <- c.read()
			}
		}(newClient(conn, r))
	}
}

func scratchWorld() *room {
	r1 := newRoom("Room 1", "You are in the first room")
	r2 := newRoom("Room 2", "You are in the second room")
	r3 := newRoom("Room 3", "You are in the third room")

	r1.exits = append(r1.exits, newExit(r2, south))
	r1.exits = append(r1.exits, newExit(r3, west))

	m := &mob{
		name:        "a test mob",
		description: "A test mob",
		room:        r1,
		roles:       []role{mobile},
	}
	r1.mobs = append(r1.mobs, m)
	mobs = append(mobs, m)

	r2.exits = append(r2.exits, newExit(r1, north))
	r3.exits = append(r3.exits, newExit(r1, east))

	r1.items = append(r1.items, newItem("an item", "An item is here"))

	return r1
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
