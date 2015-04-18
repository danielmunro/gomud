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
	tick  Event = "tick"
	pulse Event = "pulse"
)

const tickLength int = 15

type Message struct {
	client *Client
	message string
}

func (m *Message) Process() bool {
	if (m.client.mob.Delay == 0) {
		m.client.Write(m.client.mob.Act(m.message))
		return true
	}
	return false
}

type Server struct {
	clients []*Client
	listener net.Listener
	messages []*Message
	port    int
}

func NewServer(port int) *Server {
	return &Server{port: port}
}

func (s *Server) Run() {
	s.connect()
	newClientListener := make(chan *Client)
	pulseListener := make(chan Event)
	tickListener := make(chan Event)
	bufListener := make(chan *Message)
	// Listen for new clients
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				log.Fatalln(err)
			}
			newClientListener <- NewClient(conn)
		}
	}()
	// Timing
	go func() {
		nt := rand.Intn(tickLength) + tickLength
		p := 0
		for {
			time.Sleep(time.Second)
			p += 1
			pulseListener <- pulse
			if p >= nt {
				tickListener <- tick
				p = 0
				nt = rand.Intn(tickLength) + tickLength
				log.Println("Next tick in " + strconv.Itoa(nt) + " seconds")
			}
		}
	}()
	for {
		select {
		case client := <-newClientListener:
			s.clients = append(s.clients, client)
			go client.Listen(bufListener)
			log.Println("Client connected, " + strconv.Itoa(len(s.clients)) + " active clients")
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
		case message := <-bufListener:
			s.messages = append(s.messages, message)
		}
		s.processMessages()
	}
}

func (s *Server) processMessages() {
	var unprocessed []*Message
	for _, m := range s.messages {
		if (!m.Process()) {
			unprocessed = append(unprocessed, m)
		}
	}
	s.messages = unprocessed
}

func (s *Server) connect() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		log.Fatalln(err)
	} else {
		s.listener = ln
		log.Println("Listening on port :" + strconv.Itoa(s.port))
	}
}

func (s *Server) removeClient(c *Client) {
	c.conn.Close()
	for i, cl := range s.clients {
		if cl == c {
			s.clients = append(s.clients[0:i], s.clients[i+1:]...)
			log.Println("Client disconnected, " + strconv.Itoa(len(s.clients)) + " active clients")
			return
		}
	}
}
