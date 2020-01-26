package gomud

type status string

const (
	CompletedStatus status = "completed"
	ErrorStatus status = "error"
	FailedStatus status = "failed"
)

type output struct {
	input *input
	status status
	messageToRequestCreator string
	messageToTarget string
	messageToObservers string
}

func newOutput(input *input, status status, messageToRequestCreator string, messageToTarget string, messageToObservers string) *output {
	return &output{
		input,
		status,
		messageToRequestCreator,
		messageToTarget,
		messageToObservers,
	}
}

func newOutputToRequestCreator(input *input, status status, messageToRequestCreator string) *output {
	return &output{
		input,
		status,
		messageToRequestCreator,
		"",
		"",
	}
}
