package gomud

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

type Client struct {
	Conn net.Conn
	Mob  *Mob
	Buf  []string
}

func NewClient(conn net.Conn) *Client {
	c := &Client{
		Conn: conn,
		Mob:  NewMob(),
	}
	c.Mob.client = c
	c.Write("Hello World!\n")
	c.Act("look")
	c.Prompt()
	return c
}

func (c *Client) Write(line string) {
	c.Conn.Write([]byte(line))
}

func (c *Client) Act(act string) {
	c.Write(c.Mob.Act(act))
}

func (c *Client) Listen(ch chan *Client) {
	for {
		buf, _ := bufio.NewReader(c.Conn).ReadString('\n')
		c.Buf = append(c.Buf, strings.TrimSpace(buf))
		if c.Mob.Delay == 0 {
			ch <- c
		}
	}
}

func (c *Client) BufPop() string {
	b := c.Buf[0]
	c.Buf = c.Buf[1:]
	return b
}

func (c *Client) Prompt() {
	a := c.Mob.CurrentAttr
	c.Write("[" + strconv.FormatFloat(a.Hp, 'f', 0, 32) + "hp " + strconv.FormatFloat(a.Mana, 'f', 0, 32) + "m " + strconv.FormatFloat(a.Mv, 'f', 0, 32) + "mv]> ")
}

func (c *Client) FlushBuf() {
	output := false
	if c.Mob.Delay == 0 {
		for {
			if len(c.Buf) > 0 {
				b := c.BufPop()
				c.Act(b)
				output = true
			} else {
				break
			}
		}
	}
	if output {
		c.Prompt()
	}
}

func (c *Client) Pulse() {
	c.Mob.DecrementDelay()
	c.FlushBuf()
}
