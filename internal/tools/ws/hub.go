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
	Mutex      sync.Mutex
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
			h.Mutex.Lock()
			h.Clients[client] = true
			h.Mutex.Unlock()
		case client := <-h.Unregister:
			h.Mutex.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				client.Conn.Close()
			}
			h.Mutex.Unlock()
		case message := <-h.Broadcast:
			h.Mutex.Lock()
			for client := range h.Clients {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					client.Conn.Close()
					delete(h.Clients, client)
				}
			}
			h.Mutex.Unlock()

			// case message := <-h.Broadcast:
			// 	h.mutex.Lock()
			// 	for client := range h.Clients {
			// 		client.mutex.Lock()
			// 		if client.Channels[message.Channel] {
			// 			err := client.Conn.WriteJSON(message)
			// 			if err != nil {
			// 				client.Conn.Close()
			// 				delete(h.Clients, client)
			// 			}
			// 		}
			// 		client.mutex.Unlock()
			// 	}
			// 	h.mutex.Unlock()
		}
	}
}
