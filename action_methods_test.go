package gomud

import (
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
	client := newClient(&net.TCPConn{})
	gs.dummyLogin(client)

	output := gs.HandleBuffer(&Buffer{
		input: "look",
		client: client,
	})

	if output.status != CompletedStatus {
		t.Error("expected completed status")
	}
	if output.messageToRequestCreator != room1 {
		t.Error("expected message: " + output.messageToRequestCreator)
	}
}
