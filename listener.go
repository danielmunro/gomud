package gomud

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type client struct {
	conn    net.Conn
	room    *room
	message string
}

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

func newClient(conn net.Conn) *client {
	return &client{
		conn: conn,
	}
}

// Listen ...
func Listen(port int) {
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

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go func() {
			c := newClient(conn)
			c.room = r
			clients = append(clients, c)
			look(&input{client: c})
			for {
				listener <- c.read()
			}
		}()
	}
}

func scratchWorld() *room {
	r1 := newRoom("Room 1", "You are in the first room")
	r2 := newRoom("Room 2", "You are in the second room")

	r1.exits = append(r1.exits, &exit{
		room:      r2,
		direction: "south",
	})

	r2.exits = append(r2.exits, &exit{
		room:      r1,
		direction: "north",
	})

	return r1
}
