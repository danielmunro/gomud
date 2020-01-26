package gomud

import (
	"log"
	"net"
	"strconv"
	"time"
)

type Server struct {
	clients []*client
	updater chan *client
	listener net.Listener
}

func NewServer(port int) *Server {
	connection := ":"+strconv.Itoa(port)
	listener, err := net.Listen("tcp", connection)
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		updater: make(chan *client),
		listener: listener,
	}
}

func (s *Server) Listen(bufferWriter chan *Buffer) error {
	defer s.listener.Close()
	go s.readClientInput(bufferWriter)
	go s.loop()
	log.Printf("server started on %s", s.listener.Addr().String())
	for {
		client := newClient(listen(s.listener))
		go s.addClientAndListen(client)
		log.Printf("connection established from %s", client.conn.RemoteAddr().String())
	}
}

func (s *Server) addClientAndListen(c *client) {
	s.clients = append(s.clients, c)
	for {
		s.updater <- c.read()
	}
}

func (s *Server) readClientInput(bufferWriter chan *Buffer) {
	for {
		select {
		case c := <- s.updater:
			bufferWriter <- NewBuffer(c, c.message)
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
				c.writePrompt("")
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
