package gomud

import (
	"github.com/danielmunro/gomud/io"
	"net"
	"testing"
)

const room1 = `Room 1
You are in the first Room
[sw]
an item is here.
an item is here.
a test Mob is here.
`

func Test_Look_AtRoom(t *testing.T) {
	gs := NewGameService(io.NewServer(1234))
	gs.CreateFixtures()
	client := io.NewClient(&net.TCPConn{})
	gs.dummyLogin(client)

	output := gs.HandleBuffer(&io.Buffer{
		Input: "look",
		Client: client,
	})

	if output.Status != io.CompletedStatus {
		t.Error("expected completed Status")
	}
	if output.MessageToRequestCreator != room1 {
		t.Error("expected message: " + output.MessageToRequestCreator)
	}
}
