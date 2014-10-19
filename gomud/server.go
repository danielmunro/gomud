package gomud

import (
	"log"
	"net"
	"strconv"
	"time"
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
	timeKeeper := make(chan int)
	go s.ConnectionListener(ln, newClientListener)
	go s.TimeKeeper(timeKeeper)
	for {
		select {
		case client := <-newClientListener:
			go client.Listen(clientListener)
			s.clients = append(s.clients, client)
		case client := <-clientListener:
			client.FlushBuf()
		case <-timeKeeper:
			for _, cl := range s.clients {
				cl.Pulse()
			}
		}
	}
}

func (server *Server) TimeKeeper(timekeeper chan int) {
	s := time.Now().Second()
	for {
		if time.Now().Second() != s {
			s = time.Now().Second()
			timekeeper <- s
		}
	}
}

func (s *Server) ConnectionListener(ln net.Listener, ch chan *Client) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		ch <- NewClient(conn)
	}
}
