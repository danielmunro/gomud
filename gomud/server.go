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
	Tick Event = "tick"
	Pulse Event = "pulse"
)

type Server struct {
	clients  []*Client
	port     int
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
	timeKeeper := make(chan Event)
	go s.connectionListener(ln, newClientListener)
	go s.timeKeeper(timeKeeper)
	for {
		select {
		case client := <-newClientListener:
			go client.Listen(clientListener)
			s.clients = append(s.clients, client)
		case client := <-clientListener:
			client.FlushBuf()
		case event := <-timeKeeper:
			if event == Pulse {
				for _, m := range mobs {
					m.Pulse()
				}
				for _, cl := range s.clients {
					cl.Pulse()
				}
			} else if event == Tick {
				for _, m := range mobs {
					m.Tick()
				}
				for _, cl := range s.clients {
					cl.Tick()
				}
			}
		}
	}
}

func (server *Server) timeKeeper(timekeeper chan Event) {
	s := time.Now().Second()
	rand.Seed(time.Now().Unix())
	nextTick := rand.Intn(15) + 15
	pulse := 0
	log.Println("Next tick in "+strconv.Itoa(nextTick)+" seconds")
	for {
		if time.Now().Second() != s {
			s = time.Now().Second()
			pulse += 1
			timekeeper <- Pulse
			if pulse >= nextTick {
				timekeeper <- Tick
				nextTick = rand.Intn(15) + 15
				pulse = 0
				log.Println("Next tick in "+strconv.Itoa(nextTick)+" seconds")
			}
		}
	}
}

func (s *Server) connectionListener(ln net.Listener, ch chan *Client) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		ch <- NewClient(conn)
	}
}
