package domain

import (
	"sync"
)

type Pool struct {
	Register  chan *Client
	Leave     chan *Client
	Broadcast chan Message
	Clients   map[*Client]bool
	sync.Mutex
}

func NewPool() *Pool {
	return &Pool{
		Register:  make(chan *Client),
		Leave:     make(chan *Client),
		Broadcast: make(chan Message),
		Clients:   make(map[*Client]bool),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = false
		case client := <-p.Leave:
			delete(p.Clients, client)
		case message := <-p.Broadcast:
			for client := range p.Clients {
				client.send <- message
			}
		}
	}
}
