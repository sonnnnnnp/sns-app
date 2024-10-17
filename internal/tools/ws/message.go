package ws

type Message struct {
	Type    string      `json:"type,omitempty"`
	Channel string      `json:"channel,omitempty"`
	Params  interface{} `json:"params,omitempty"`
	Body    interface{} `json:"body,omitempty"`
}
