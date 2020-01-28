package gomud

import (
	"github.com/danielmunro/gomud/io"
	"net"
	"testing"
)

type Test struct {
	gameService *GameService
	client *io.Client
	t *testing.T
}

func NewTest(t *testing.T) *Test {
	server, _ := io.NewServer(1234)
	gs := NewGameService(server)
	gs.CreateFixtures()
	client := io.NewClient(&net.TCPConn{})
	gs.dummyLogin(client)
	return &Test{
		gameService: gs,
		client: client,
		t: t,
	}
}

func (t *Test) GetOutputFromInput(input string) *io.Output {
	return t.gameService.HandleBuffer(&io.Buffer{
		Input: input,
		Client: t.client,
	})
}

func (t *Test) Expect(condition bool, fail string) {
	if !condition {
		t.t.Fatal(fail)
	}
}
