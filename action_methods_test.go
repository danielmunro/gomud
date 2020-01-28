package gomud

import (
	"github.com/danielmunro/gomud/io"
	"testing"
)

const room1 = `Room 1
You are in the first Room
[sw]
an item is here.
an item is here.
a test Mob is here.
`

const room2 = `Room 3
You are in the third Room
[e]
`

func Test_Look_AtRoom(t *testing.T) {
	// setup
	test := NewTest(t)

	// when
	output := test.GetOutputFromInput("look")

	// then
	test.Expect(output.Status == io.CompletedStatus, "expected completed Status")
	test.Expect(output.MessageToRequestCreator == room1, "expected message: " + output.MessageToRequestCreator)
}

func Test_Look_AfterMove(t *testing.T) {
	// setup
	test := NewTest(t)

	// when
	output := test.GetOutputFromInput("w")

	// then
	test.Expect(output.MessageToRequestCreator == room2, "expected output to be room2")
}
