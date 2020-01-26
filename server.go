package gomud

import (
	"github.com/danielmunro/gomud/io"
	"log"
	"net"
	"strconv"
	"time"
)

type Server struct {
	clients []*io.Client
	updater chan *io.Client
	listener net.Listener
}

func NewServer(port int) *Server {
	connection := ":"+strconv.Itoa(port)
	listener, err := net.Listen("tcp", connection)
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		updater: make(chan *io.Client),
		listener: listener,
	}
}

func (s *Server) Listen(bufferWriter chan *io.Buffer) error {
	defer s.listener.Close()
	go s.readClientInput(bufferWriter)
	go s.loop()
	log.Printf("server started on %s", s.listener.Addr().String())
	for {
		client := io.NewClient(listen(s.listener))
		go s.addClientAndListen(client)
		log.Printf("connection established from %s", client.String())
	}
}

func (s *Server) addClientAndListen(c *io.Client) {
	s.clients = append(s.clients, c)
	for {
		s.updater <- c.Read()
	}
}

func (s *Server) readClientInput(bufferWriter chan *io.Buffer) {
	for {
		select {
		case c := <- s.updater:
			bufferWriter <- io.NewBuffer(c, c.Message)
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
