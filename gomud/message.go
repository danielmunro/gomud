package gomud

type Message struct {
	client *Client
	data string
}

func (m *Message) process() bool {
	if (m.client.mob.Delay == 0) {
		m.client.write(m.client.mob.Act(m.data))
		return true
	}
	return false
}

func NewMessage(client *Client, data string) *Message {
	return &Message{client: client, data: data}
}
