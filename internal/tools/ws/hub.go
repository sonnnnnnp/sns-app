package ws

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Hub struct {
	Upgrader       *websocket.Upgrader
	Clients        map[*Client]bool
	RegisterChan   chan *Client
	UnregisterChan chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Upgrader:       GetUpgrader(nil),
		Clients:        make(map[*Client]bool),
		RegisterChan:   make(chan *Client),
		UnregisterChan: make(chan *Client),
	}
}

func (h *Hub) ContainsUser(uIDs []uuid.UUID, uID uuid.UUID) bool {
	for _, u := range uIDs {
		if u == uID {
			return true
		}
	}
	return false
}

func (h *Hub) Listen() {
	for {
		select {
		case client := <-h.RegisterChan:
			h.Clients[client] = true
		case client := <-h.UnregisterChan:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				client.Conn.Close()
			}
		}
	}
}
