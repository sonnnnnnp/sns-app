package ws

type Room struct {
	Clients []*Client
}

func (r *Room) Add(client *Client) {
	r.Clients = append(r.Clients, client)
}

func (r *Room) Get(client *Client) []Client {
	var cs []Client

	for _, c := range r.Clients {
		cs = append(cs, *c)
	}

	return cs
}

func (r *Room) Publish(msg []byte) {
	for _, client := range r.Clients {
		client.Send(msg)
	}
}
