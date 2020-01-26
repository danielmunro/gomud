package gomud

import "github.com/danielmunro/gomud/io"

type status string

const (
	CompletedStatus status = "completed"
	ErrorStatus status = "error"
	FailedStatus status = "failed"
)

type output struct {
	buffer *io.Buffer
	status status
	messageToRequestCreator string
	messageToTarget string
	messageToObservers string
}

func newOutput(buffer *io.Buffer, status status, messageToRequestCreator string, messageToTarget string, messageToObservers string) *output {
	return &output{
		buffer,
		status,
		messageToRequestCreator,
		messageToTarget,
		messageToObservers,
	}
}

func newOutputToRequestCreator(buffer *io.Buffer, status status, messageToRequestCreator string) *output {
	return &output{
		buffer,
		status,
		messageToRequestCreator,
		"",
		"",
	}
}
