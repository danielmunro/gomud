package io

type Status string

const (
	CompletedStatus Status = "completed"
	ErrorStatus     Status = "error"
	FailedStatus    Status = "failed"
)

type Output struct {
	Buffer                  *Buffer
	Status                  Status
	MessageToRequestCreator string
	MessageToTarget         string
	MessageToObservers      string
}

func NewOutput(buffer *Buffer, status Status, messageToRequestCreator string, messageToTarget string, messageToObservers string) *Output {
	return &Output{
		buffer,
		status,
		messageToRequestCreator,
		messageToTarget,
		messageToObservers,
	}
}

func NewOutputToRequestCreator(buffer *Buffer, status Status, messageToRequestCreator string) *Output {
	return &Output{
		buffer,
		status,
		messageToRequestCreator,
		"",
		"",
	}
}
