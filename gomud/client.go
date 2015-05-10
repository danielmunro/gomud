package gomud

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

/*
	Client encapsulates the data necessary to run a single
	connection to the server.
*/
type Client struct {
	//conn is the connection to the client
	conn net.Conn
	//mob is a pointer to a mobile that is controlled by the user during their session.
	mob *Mob
	//buf holds messages from the Client before they are sent onward.
	buf []string
	//server is a connection to the Server that is overseeing the client's session
	server *Server
}

/*
	NewClient creates a new Client struct for the given Connection.
*/
func NewClient(conn net.Conn) *Client {
	//creates a new client with the connection and a new Mob
	c := &Client{
		conn: conn,
		mob:  NewMob(),
	}
	//sets the client of the client's mob to be the new client
	c.mob.client = c
	//simulates the user input "look"
	c.Write(c.mob.Act("look"))
	//prompts the user for input
	c.prompt()
	return c
}

/*
	Write writes a string to the Client's net.Conn connection.
*/
func (c *Client) Write(line string) {
	c.conn.Write([]byte(line))
}

/*
	Listen waits for input on the Client's net.Conn connection and sends it to the
	provided bufListener channel.
*/
func (c *Client) Listen(bufListener chan<- *Message) {
	for {
		buf, _ := bufio.NewReader(c.conn).ReadString('\n')
		bufListener <- NewMessage(c, strings.TrimSpace(buf))
	}
}

/*
	Pulse notifies the Client's mob and prompts the Client.
*/
func (c *Client) Pulse() {
	if c.mob.target != nil {
		c.mob.Notify(c.mob.target.ShortName + " " + c.mob.target.Status() + ".\n\n")
		c.prompt()
	}
}

/*
	Tick writes a newline to the Client and prompts the client.
*/
func (c *Client) Tick() {
	c.Write("\n")
	c.prompt()
}

/*
	bufPop returns the first string in the Client's buf string array.
*/
func (c *Client) bufPop() string {
	b := c.buf[0]
	c.buf = c.buf[1:]
	return b
}

/*
	prompt displays information about the Client mob's current attributes.
*/
func (c *Client) prompt() {
	a := c.mob.CurrentAttr
	c.write("[" + strconv.FormatFloat(a.Vitals.Hp, 'f', 0, 32) + "hp " + strconv.FormatFloat(a.Vitals.Mana, 'f', 0, 32) + "m " + strconv.FormatFloat(a.Vitals.Mv, 'f', 0, 32) + "mv]> ")
}
