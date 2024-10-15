package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Upgrader   *websocket.Upgrader
	Clients    map[*Client]bool
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
	mutex      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Upgrader:   GetUpgrader(nil),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Start() {
	for {
		select {
		case client := <-h.Register:
			h.mutex.Lock()
			h.Clients[client] = true
			h.mutex.Unlock()
		case client := <-h.Unregister:
			h.mutex.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				client.Conn.Close()
			}
			h.mutex.Unlock()
		case message := <-h.Broadcast:
			h.mutex.Lock()
			for client := range h.Clients {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					client.Conn.Close()
					delete(h.Clients, client)
				}
			}
			h.mutex.Unlock()
		}
	}
}
