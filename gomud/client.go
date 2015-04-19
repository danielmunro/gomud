package gomud

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

type Client struct {
	conn   net.Conn
	mob    *Mob
	buf    []string
	server *Server
}

func NewClient(conn net.Conn) *Client {
	c := &Client{
		conn: conn,
		mob:  NewMob(),
	}
	c.mob.client = c
	c.Write(c.mob.Act("look"))
	c.prompt()
	return c
}

func (c *Client) Write(line string) {
	c.conn.Write([]byte(line))
}

func (c *Client) Listen(bufListener chan<-*Message) {
	for {
		buf, _ := bufio.NewReader(c.conn).ReadString('\n')
		bufListener <- &Message{
			client: c,
			message: strings.TrimSpace(buf),
		}
	}
}

func (c *Client) Pulse() {
	if c.mob.target != nil {
		c.mob.Notify(c.mob.target.ShortName + " " + c.mob.target.Status() + ".\n\n")
		c.prompt()
	}
}

func (c *Client) Tick() {
	c.Write("\n")
	c.prompt()
}

func (c *Client) bufPop() string {
	b := c.buf[0]
	c.buf = c.buf[1:]
	return b
}

func (c *Client) prompt() {
	a := c.mob.CurrentAttr
	c.Write("[" + strconv.FormatFloat(a.Hp, 'f', 0, 32) + "hp " + strconv.FormatFloat(a.Mana, 'f', 0, 32) + "m " + strconv.FormatFloat(a.Mv, 'f', 0, 32) + "mv]> ")
}
