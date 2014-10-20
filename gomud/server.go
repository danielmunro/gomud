package gomud

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type Event string

const (
	Tick  Event = "tick"
	Pulse Event = "pulse"
)

type Server struct {
	clients []*Client
	port    int
}

func NewServer() *Server {
	return &Server{port: 8080}
}

func (s *Server) Run() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Listening on port :" + strconv.Itoa(s.port))
	newClientListener := make(chan *Client)
	clientListener := make(chan *Client)
	pulseListener := make(chan Event)
	tickListener := make(chan Event)
	go connectionListener(ln, newClientListener)
	go timeKeeper(pulseListener, tickListener)
	for {
		select {
		case client := <-newClientListener:
			s.addClient(client, clientListener)
		case client := <-clientListener:
			client.FlushBuf()
		case <-pulseListener:
			for _, m := range mobs {
				m.Pulse()
			}
			for _, cl := range s.clients {
				cl.Pulse()
			}
		case <-tickListener:
			for _, m := range mobs {
				m.Tick()
			}
			for _, cl := range s.clients {
				cl.Tick()
			}
		}
	}
}

func (s *Server) addClient(c *Client, listener chan *Client) {
	go c.Listen(listener)
	s.clients = append(s.clients, c)
	c.server = s
	log.Println("Client connected, " + strconv.Itoa(len(s.clients)) + " active clients")
}

func (s *Server) removeClient(c *Client) {
	for i, cl := range s.clients {
		if cl == c {
			s.clients = append(s.clients[0:i], s.clients[i+1:]...)
			log.Println("Client disconnected, " + strconv.Itoa(len(s.clients)) + " active clients")
			return
		}
	}
}

func timeKeeper(pulseListener chan Event, tickListener chan Event) {
	t := time.Now().Second()
	nt := nextTick()
	pulse := 0
	for {
		if time.Now().Second() != t {
			t = time.Now().Second()
			pulse += 1
			pulseListener <- Pulse
			if pulse >= nt {
				tickListener <- Tick
				pulse = 0
				nt = nextTick()
			}
		}
	}
}

func connectionListener(ln net.Listener, ch chan *Client) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		ch <- NewClient(conn)
	}
}

func nextTick() int {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(15) + 15
	log.Println("Next tick in " + strconv.Itoa(n) + " seconds")
	return n
}
