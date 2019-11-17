package domain

import (
	"github.com/gorilla/websocket"
	"github.com/terawork-com/message-socket-service/pkg/log"
)

type Client struct {
	ID         string
	Conn       *websocket.Conn
	Pool       *Pool
	send       chan []byte
	disconnect chan int8
}

func (c *Client) Read() {
	defer func() {
		c.closeConnection()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Error(err)
			return
		}
		c.processMessage(msg)
	}
}

func (c *Client) Write() {
	defer func() {
		c.closeConnection()
	}()

	for msg := range c.send {
		err := c.Conn.WriteJSON(msg)
		if err != nil {
			log.Error(err)
		}
		log.Info("User left")
	}
}

func (c *Client) closeConnection() {
	err := c.Conn.Close()
	if err != nil {
		log.Errorf("Error closing client connection: %v", err)
	}
}

func (c *Client) processMessage(m Message) (msg Message, hasMessage bool) {
	log.Info("Message received")
	return
}
