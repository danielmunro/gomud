package gomud

// Message encapsulates a string of text and an associated client.
type Message struct {
	client *Client
	data   string
}

// Process executes the message's string if the Client's Mob can
// act and returns whether or not it was able to perform the
// string's action.
func (m *Message) Process() bool {
	if m.client.mob.Delay == 0 {
		m.client.write(m.client.mob.Act(m.data))
		return true
	}
	return false
}

// NewMessage creates and returns a new Message struct containing the data
// string and associated with the given client.
func NewMessage(client *Client, data string) *Message {
	return &Message{client: client, data: data}
}
