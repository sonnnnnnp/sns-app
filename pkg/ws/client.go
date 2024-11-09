package ws

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	UserID   uuid.UUID
	Conn     *websocket.Conn
	Channels map[string]bool
}

func NewClient(uID uuid.UUID, conn *websocket.Conn) *Client {
	return &Client{
		UserID:   uID,
		Conn:     conn,
		Channels: make(map[string]bool),
	}
}

func (c *Client) Send(msg Message) error {
	return c.Conn.WriteJSON(msg)
}

func (c *Client) Subscribe(channel string) error {
	if c.Channels[channel] {
		return errors.New("channel already subscribed")
	}

	c.Channels[channel] = true

	return nil
}

func (c *Client) Unsubscribe(channel string) error {
	if !c.Channels[channel] {
		return errors.New("channel not subscribed")
	}

	delete(c.Channels, channel)

	return nil
}

func (c *Client) Ping() {
	ticker := time.NewTicker(time.Second * 5)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for range ticker.C {
		c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		if err := c.Send(Message{Type: "ping"}); err != nil {
			return
		}
	}
}
