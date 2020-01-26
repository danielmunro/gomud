package gomud

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"net"
	"strings"
)

type client struct {
	conn    net.Conn
	mob     *mob
	message string
	id      string
}

func (c *client) read() *client {
	c.message, _ = bufio.NewReader(c.conn).ReadString('\n')
	c.message = strings.Trim(c.message, "\r\n")

	return c
}

func (c *client) writePrompt(m string) {
	c.conn.Write([]byte(fmt.Sprintf("%s\n--> ", m)))
}

func (c *client) String() string {
	return c.conn.RemoteAddr().String()
}

func newClient(c net.Conn) *client {
	cl := &client{
		conn: c,
		id: uuid.New().String(),
	}

	return cl
}
