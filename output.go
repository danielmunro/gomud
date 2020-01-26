package gomud

import "github.com/danielmunro/gomud/io"

type status string

const (
	CompletedStatus status = "completed"
	ErrorStatus status = "error"
	FailedStatus status = "failed"
)

type output struct {
	input *io.Input
	status status
	messageToRequestCreator string
	messageToTarget string
	messageToObservers string
}

func newOutput(input *io.Input, status status, messageToRequestCreator string, messageToTarget string, messageToObservers string) *output {
	return &output{
		input,
		status,
		messageToRequestCreator,
		messageToTarget,
		messageToObservers,
	}
}

func newOutputToRequestCreator(input *io.Input, status status, messageToRequestCreator string) *output {
	return &output{
		input,
		status,
		messageToRequestCreator,
		"",
		"",
	}
}
