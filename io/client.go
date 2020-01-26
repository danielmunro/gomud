package io

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"net"
	"strings"
)

type Client struct {
	conn    net.Conn
	//mob     *gomud.Mob
	Message string
	id      string
}

func (c *Client) Read() *Client {
	c.Message, _ = bufio.NewReader(c.conn).ReadString('\n')
	c.Message = strings.Trim(c.Message, "\r\n")

	return c
}

func (c *Client) WritePrompt(m string) {
	c.conn.Write([]byte(fmt.Sprintf("%s\n--> ", m)))
}

func (c *Client) String() string {
	return c.conn.RemoteAddr().String()
}

func NewClient(c net.Conn) *Client {
	cl := &Client{
		conn: c,
		id: uuid.New().String(),
	}

	return cl
}
