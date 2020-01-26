package gomud

import (
	"github.com/danielmunro/gomud/io"
	"net"
	"testing"
)

const room1 = `Room 1
You are in the first room
[sw]
an item is here.
an item is here.
`

func Test_Look_AtRoom(t *testing.T) {
	gs := NewGameService(NewServer(1234))
	gs.CreateFixtures()
	client := io.NewClient(&net.TCPConn{})
	gs.dummyLogin(client)

	output := gs.HandleBuffer(&io.Buffer{
		Input: "look",
		Client: client,
	})

	if output.status != CompletedStatus {
		t.Error("expected completed status")
	}
	if output.messageToRequestCreator != room1 {
		t.Error("expected message: " + output.messageToRequestCreator)
	}
}
