package gomud

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

// Client encapsulates the data necessary to run a single
// connection to the server.
// conn is the connection to the client
// mob is a pointer to a mobile that is controlled by the user during their session.
// buf holds messages from the Client before they are sent onward.
// server is a connection to the Server that is overseeing the client's session
type Client struct {
	conn net.Conn
	mob *Mob
	buf []string
	server *Server
}

// NewClient creates a new Client struct for the given Connection.
func NewClient(conn net.Conn) *Client {
	//creates a new client with the connection and a new Mob
	c := &Client{
		conn: conn,
		mob:  NewMob(),
	}
	//sets the client of the client's mob to be the new client
	c.mob.client = c
	//simulates the user input "look"
	c.write(c.mob.Act("look"))
	//prompts the user for input
	c.prompt()
	return c
}

// write writes a string to the Client's net.Conn connection.
func (c *Client) write(line string) {
	c.conn.Write([]byte(line))
}

// listen waits for input on the Client's net.Conn connection and sends it to
// the provided bufListener channel.
func (c *Client) listen(bufListener chan<- *Message) {
	for {
		buf, _ := bufio.NewReader(c.conn).ReadString('\n')
		bufListener <- NewMessage(c, strings.TrimSpace(buf))
	}
}

// pulse notifies the Client's mob and prompts the Client.
func (c *Client) pulse() {
	if c.mob.target != nil {
		c.mob.Notify(c.mob.target.ShortName + " " + c.mob.target.Status() + ".\n\n")
		c.prompt()
	}
}

// tick writes a newline to the Client and prompts the client.
func (c *Client) tick() {
	c.write("\n")
	c.prompt()
}

// bufPop returns the first string in the Client's buf string array.
func (c *Client) bufPop() string {
	b := c.buf[0]
	c.buf = c.buf[1:]
	return b
}

// prompt displays information about the Client mob's current attributes.
func (c *Client) prompt() {
	a := c.mob.CurrentAttr
	c.write("[" + strconv.FormatFloat(a.Vitals.Hp, 'f', 0, 32) + "hp " + strconv.FormatFloat(a.Vitals.Mana, 'f', 0, 32) + "m " + strconv.FormatFloat(a.Vitals.Mv, 'f', 0, 32) + "mv]> ")
}
