package gomud

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

type Client struct {
	Conn   net.Conn
	Mob    *Mob
	Buf    []string
	server *Server
}

func NewClient(conn net.Conn) *Client {
	c := &Client{
		Conn: conn,
		Mob:  NewMob(),
	}
	c.Mob.client = c
	c.Write(c.Mob.Act("look"))
	c.prompt()
	return c
}

func (c *Client) Write(line string) {
	c.Conn.Write([]byte(line))
}

func (c *Client) Listen(ch chan *Client) {
	for {
		buf, _ := bufio.NewReader(c.Conn).ReadString('\n')
		c.Buf = append(c.Buf, strings.TrimSpace(buf))
		ch <- c
	}
}

func (c *Client) FlushBuf() {
	output := false
	if c.Mob.Delay == 0 {
		for {
			if len(c.Buf) > 0 {
				b := c.bufPop()
				c.Write(c.Mob.Act(b))
				output = true
			} else {
				break
			}
		}
	}
	if output {
		c.prompt()
	}
}

func (c *Client) Pulse() {
	c.FlushBuf()
}

func (c *Client) Tick() {
	c.Write("\n")
	c.prompt()
}

func (c *Client) bufPop() string {
	b := c.Buf[0]
	c.Buf = c.Buf[1:]
	return b
}

func (c *Client) prompt() {
	a := c.Mob.CurrentAttr
	c.Write("[" + strconv.FormatFloat(a.Hp, 'f', 0, 32) + "hp " + strconv.FormatFloat(a.Mana, 'f', 0, 32) + "m " + strconv.FormatFloat(a.Mv, 'f', 0, 32) + "mv]> ")
}
