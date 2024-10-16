package ws

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	UserID   uuid.UUID
	Conn     *websocket.Conn
	Channels map[string]bool
	mutex    sync.Mutex
}

func NewClient(uID uuid.UUID, conn *websocket.Conn) *Client {
	return &Client{
		UserID:   uID,
		Conn:     conn,
		Channels: make(map[string]bool),
	}
}

func (c *Client) Send(msg Message) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.Conn.WriteJSON(msg)
}

func (c *Client) Subscribe(channel string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Channels[channel] {
		return errors.New("channel already subscribed")
	}

	c.Channels[channel] = true

	return nil
}

func (c *Client) Unsubscribe(channel string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if !c.Channels[channel] {
		return errors.New("channel not subscribed")
	}

	delete(c.Channels, channel)

	return nil
}
