package io

import (
	"log"
	"net"
	"strconv"
	"time"
)

type Server struct {
	clients []*Client
	updater chan *Client
	listener net.Listener
}

func NewServer(port int) (*Server, error) {
	connection := ":"+strconv.Itoa(port)
	listener, err := net.Listen("tcp", connection)
	if err != nil {
		return nil, err
	}
	return &Server{
		updater: make(chan *Client),
		listener: listener,
	}, nil
}

func (s *Server) Listen(bufferWriter chan *Buffer) error {
	defer s.listener.Close()
	go s.readClientInput(bufferWriter)
	go s.loop()
	log.Printf("server started on %s", s.listener.Addr().String())
	for {
		client := NewClient(listen(s.listener))
		go s.addClientAndListen(client)
		log.Printf("connection established from %s", client.String())
	}
}

func (s *Server) addClientAndListen(c *Client) {
	s.clients = append(s.clients, c)
	for {
		s.updater <- c.Read()
	}
}

func (s *Server) readClientInput(bufferWriter chan *Buffer) {
	for {
		select {
		case c := <- s.updater:
			bufferWriter <- NewBuffer(c, c.Message)
		}
	}
}

func (s *Server) loop() {
	const (
		tickLen time.Duration = 15
	)
	pulse := time.NewTicker(time.Second)
	tick := time.NewTicker(time.Second * tickLen)
	for {
		select {
		case <-pulse.C:

			break
		case <-tick.C:
			for _, c := range s.clients {
				c.WritePrompt("")
			}
			break
		}
	}
}

func listen(l net.Listener) net.Conn {
	conn, err := l.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}
