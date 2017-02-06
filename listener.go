package gomud

import (
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type Listener struct {
	mobs    []*mob
	clients []*client
	updater chan *client
}

func NewListener() *Listener {
	return &Listener{}
}

func (l *Listener) Listen(port int) error {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	defer ln.Close()
	l.updater = make(chan *client)
	go l.inputReader()
	go l.timing()

	for {
		go l.addClient(newClient(newConnection(ln), startRoom(l)))
	}
}

func (l *Listener) addClient(c *client) {
	l.clients = append(l.clients, c)
	l.mobs = append(l.mobs, c.mob)
	newAction(c.mob, "look")
	for {
		l.updater <- c.read()
	}
}

func (l *Listener) inputReader() {
	for {
		select {
		case c := <-l.updater:
			newActionWithInput(newInput(c, strings.Split(c.message, " ")))
		}
	}
}

func (l *Listener) timing() {
	const (
		tickLen time.Duration = 15
	)
	pulse := time.NewTicker(time.Second)
	tick := time.NewTicker(time.Second * tickLen)
	for {
		select {
		case <-pulse.C:
			for _, m := range l.mobs {
				if d20() == 1 {
					roleCheck(m)
				}
			}
			break
		case <-tick.C:
			for _, m := range l.mobs {
				regen(m)
			}
			for _, c := range l.clients {
				c.writePrompt("")
			}
			break
		}
	}
}

func newConnection(l net.Listener) net.Conn {
	conn, err := l.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}

func roleCheck(m *mob) {
	for _, r := range m.roles {
		if r == mobile {
			m.roam()
		} else if r == scavenger {
			m.scavenge()
		}
	}
}

func regen(m *mob) {

}
