package gomud

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn    net.Conn
	mob     *mob
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

func newClient(c net.Conn, r *room) *client {
	m := &mob{
		name:        "a mob",
		description: "a mob",
		room:        r,
	}
	cl := &client{
		conn: c,
		mob:  m,
	}
	m.client = cl
	m.room.mobs = append(m.room.mobs, m)

	return cl
}
