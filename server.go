package gomud

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// tickLength is the number of <unit of time>s between each
// step in the MUD world.
const tickLength int64 = 15

// Server contains all of the data needed to run a server.
// Attributes:
// clients - an array of pointers to Client structs, used
// 			to keep track of the connected users.
// listener - a Listener from the net package, used to communicate
// 			over the specified port.
// messages - an array of pointers to Message structs, used
// 			to old outgoing messages from the server
// port - an int tracking the port number on which the server
// 			is listening
// nextTick - an int64 tracking when the next "tick" of the world
// 			should occur in unix time.
// lastTick - an int64 tracking when the last world "tick" occurred
// 			in unix time.
type Server struct {
	clients   []*Client
	listener  net.Listener
	messages  []*Message
	port      int
	nextTick  int64
	lastPulse int64
}

// NewServer creates and returns a new server listening on the given port.
func NewServer(port int) *Server {
	return &Server{port: port}
}

// Run causes the server to begin listening for client connections on its port.
func (s *Server) Run() {
	s.connect()
	newClientListener := make(chan *Client)
	bufListener := make(chan *Message)
	go s.newClientListener(newClientListener)
	for {
		select {
		case client := <-newClientListener:
			client.server = s
			s.clients = append(s.clients, client)
			go client.listen(bufListener)
			log.Println("Client connected, " + strconv.Itoa(len(s.clients)) + " active clients")
		case message := <-bufListener:
			s.messages = append(s.messages, message)
		default:
			s.timing()
			s.processMessages()
		}
	}
}

// newClientListener begins listening for new client connections and sends
// them back along the newClientListener channel if their connections are
// successful.
func (s *Server) newClientListener(newClientListener chan<- *Client) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		newClientListener <- NewClient(conn)
	}
}

// timing causes the server to execute its timing logic. This calls the
// Pulse() method on mobs and clients for every tick.
func (s *Server) timing() {
	t := time.Now().Unix()
	if t > s.lastPulse {
		s.lastPulse = t
		for _, m := range mobs {
			m.Pulse()
		}
		for _, cl := range s.clients {
			cl.pulse()
		}
	}
	if t > s.nextTick {
		s.nextTick = t + rand.Int63n(tickLength) + tickLength
		for _, m := range mobs {
			m.Tick()
		}
		for _, cl := range s.clients {
			cl.tick()
		}
	}
}

// processMessages iterates across the messages array and removes any
// messages that it was able to process successfully.
func (s *Server) processMessages() {
	for i, m := range s.messages {
		if m.Process() {
			s.messages = append(s.messages[0:i], s.messages[i+1:]...)
		}
	}
}

// connect causes the server to begin listening on its port.
func (s *Server) connect() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		log.Fatalln(err)
	} else {
		s.listener = ln
		log.Println("Listening on port :" + strconv.Itoa(s.port))
	}
}

// removeClient closes the connection to a given client and removes them
// from the clients array.
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
