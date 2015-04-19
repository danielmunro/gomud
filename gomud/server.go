package gomud

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

const tickLength int64 = 15

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
	nextTick int64
	lastPulse int64
}

func NewServer(port int) *Server {
	return &Server{port: port}
}

func (s *Server) Run() {
	s.connect()
	newClientListener := make(chan *Client)
	bufListener := make(chan *Message)
	go s.newClientListener(newClientListener)
	for {
		select {
		case client := <-newClientListener:
			s.clients = append(s.clients, client)
			go client.Listen(bufListener)
			log.Println("Client connected, " + strconv.Itoa(len(s.clients)) + " active clients")
		case message := <-bufListener:
			s.messages = append(s.messages, message)
		default:
			s.timing()
			s.processMessages()
		}
	}
}

func (s *Server) newClientListener(newClientListener chan<-*Client) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		newClientListener <- NewClient(conn)
	}
}

func (s *Server) timing() {
	t := time.Now().Unix()
	if t > s.lastPulse {
		s.lastPulse = t
		for _, m := range mobs {
			m.Pulse()
		}
		for _, cl := range s.clients {
			cl.Pulse()
		}
	}
	if t > s.nextTick {
		s.nextTick = t + rand.Int63n(tickLength) + tickLength
		for _, m := range mobs {
			m.Tick()
		}
		for _, cl := range s.clients {
			cl.Tick()
		}
	}
}

func (s *Server) processMessages() {
	for i, m := range s.messages {
		if (m.Process()) {
			s.messages = append(s.messages[0:i], s.messages[i+1:]...)
		}
	}
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
