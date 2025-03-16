package model

import (
	"net"
	"sync"
)

type Client struct {
	ID       string
	Conn     net.Conn
	Username string
}

type Message struct {
	Content   string
	Sender    string
	Timestamp string
}

type Model struct {
	clients    map[string]*Client
	clientsMux sync.RWMutex
}

func NewModel() *Model {
	return &Model{
		clients: make(map[string]*Client),
	}
}

func (m *Model) AddClient(client *Client) {
	m.clientsMux.Lock()
	defer m.clientsMux.Unlock()
	m.clients[client.ID] = client
}

func (m *Model) RemoveClient(clientID string) {
	m.clientsMux.Lock()
	defer m.clientsMux.Unlock()
	delete(m.clients, clientID)
}

func (m *Model) GetAllClients() []*Client {
	m.clientsMux.RLock()
	defer m.clientsMux.RUnlock()

	clients := make([]*Client, 0, len(m.clients))
	for _, client := range m.clients {
		clients = append(clients, client)
	}
	return clients
}
