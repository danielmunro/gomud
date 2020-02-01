package gomud

import (
	"github.com/danielmunro/gomud/io"
	"testing"
)

const room1 = `Room 1
You are in the first Room
[sw]
a baseball cap is here.
a cowboy hat is here.
an item is here.
an item is here.
a test Mob is here.
`

const room2 = `Room 3
You are in the third Room
[e]
a merchant is here.
`

func Test_Look_AtRoom(t *testing.T) {
	// setup
	test := NewTest(t)

	// when
	output := test.GetOutputFromInput("look")

	// then
	test.Expect(output.Status == io.CompletedStatus, "expected completed Status")
	test.Expect(output.MessageToRequestCreator == room1, "expected message: "+output.MessageToRequestCreator)
}

func Test_Look_AfterMove(t *testing.T) {
	// setup
	test := NewTest(t)

	// when
	output := test.GetOutputFromInput("w")

	// then
	test.Expect(output.MessageToRequestCreator == room2, "expected: " + output.MessageToRequestCreator)
}

func Test_MustBe_StandingTo_Move(t *testing.T) {
	// setup
	test := NewTest(t)

	// given
	test.GetOutputFromInput("sit")

	// when
	output := test.GetOutputFromInput("w")

	// then
	test.Expect(output.Status == io.ErrorStatus, "expected error")
	test.Expect(output.MessageToRequestCreator == "you must be standing to do that.", "requires standing to move")
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
	test.Expect(output.MessageToRequestCreator == "You pick up an item and put it in your inventory.", "should be able to pick up an item")
	test.Expect(output.MessageToObservers == "tester mctesterson picks up an item and puts it in their inventory.", "should be able to pick up an item")
}

func Test_Drop_Item(t *testing.T) {
	// setup
	test := NewTest(t)

	// given
	test.GetOutputFromInput("get item")

	// when
	output := test.GetOutputFromInput("drop item")

	// then
	test.Expect(output.Status == io.CompletedStatus, "drop request should be completed")
	test.Expect(output.MessageToRequestCreator == "You drop an item.", "should be able to drop an item")
	test.Expect(output.MessageToObservers == "tester mctesterson drops an item.", "should be able to drop an item")
}

func Test_Wear_Item(t *testing.T) {
	// setup
	test := NewTest(t)

	// given
	test.GetOutputFromInput("get hat")

	// when
	output := test.GetOutputFromInput("wear hat")

	// then
	test.Expect(output.MessageToRequestCreator == "You wear a cowboy hat.", output.MessageToRequestCreator)
}

func Test_RemoveAnd_Wear_Item(t *testing.T) {
	// setup
	test := NewTest(t)

	// given
	test.GetOutputFromInput("get hat")
	test.GetOutputFromInput("get cap")
	test.GetOutputFromInput("wear cap")

	// when
	output := test.GetOutputFromInput("wear hat")

	// then
	test.Expect(output.MessageToRequestCreator == "You remove a baseball cap and put it in your inventory. You wear a cowboy hat.", output.MessageToRequestCreator)
	test.Expect(output.MessageToObservers == "tester mctesterson removes a baseball cap and puts it in their inventory. They wear a cowboy hat.", output.MessageToObservers)
}

func Test_Kill_Mob(t *testing.T) {
	// setup
	test := NewTest(t)

	// when
	output := test.GetOutputFromInput("kill test")

	// then
	test.Expect(output.MessageToRequestCreator == "You scream and attack a test Mob!", output.MessageToRequestCreator)
	test.Expect(output.MessageToTarget == "tester mctesterson screams and attacks you!", output.MessageToTarget)
	test.Expect(output.MessageToObservers == "tester mctesterson screams and attacks a test Mob!", output.MessageToObservers)

	// and
	test.Expect(len(test.gameService.mobService.fights) == 1, "should create a fight")
}
