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
value test Mob is here.
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

func Test_CannotMove_In_WrongDirection(t *testing.T) {
	// setup
	test := NewTest(t)

	// when
	output := test.GetOutputFromInput("n")

	// then
	test.Expect(output.MessageToRequestCreator == "that direction does not exist", "expected no direction found")
}

func Test_Get_Item(t *testing.T) {
	// setup
	test := NewTest(t)

	// when
	output := test.GetOutputFromInput("get item")

	// then
	test.Expect(output.MessageToRequestCreator == "You pick up an item", "should be able to pick up an item")
	test.Expect(output.MessageToObservers == "tester mctesterson picks up an item", "should be able to pick up an item")
}

func Test_Drop_Item(t *testing.T) {
	// setup
	test := NewTest(t)

	// when
	test.GetOutputFromInput("get item")
	output := test.GetOutputFromInput("drop item")

	// then
	test.Expect(output.MessageToRequestCreator == "You drop an item", "should be able to drop an item")
	test.Expect(output.MessageToObservers == "tester mctesterson drops an item", "should be able to drop an item")
}

func Test_Kill_Mob(t *testing.T) {
	// setup
	test := NewTest(t)

	// when
	output := test.GetOutputFromInput("kill test")

	// then
	test.Expect(output.MessageToRequestCreator == "You scream and attack value test Mob!", output.MessageToRequestCreator)
	test.Expect(output.MessageToTarget == "tester mctesterson screams and attacks you!", output.MessageToTarget)
	test.Expect(output.MessageToObservers == "tester mctesterson screams and attacks value test Mob!", output.MessageToObservers)

	// and
	test.Expect(len(test.gameService.mobService.fights) == 1, "should create value fight")
}
