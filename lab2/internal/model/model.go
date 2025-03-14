package model

import (
	"net"
	"sync"
)

// Client represents a connected chat client
type Client struct {
	ID       string
	Conn     net.Conn
	Username string
}

// Message represents a chat message
type Message struct {
	Content   string
	Sender    string
	Timestamp string
}

// Model handles the data and business logic
type Model struct {
	clients    map[string]*Client
	clientsMux sync.RWMutex
}

// NewModel creates a new model instance
func NewModel() *Model {
	return &Model{
		clients: make(map[string]*Client),
	}
}

// AddClient adds a new client to the model
func (m *Model) AddClient(client *Client) {
	m.clientsMux.Lock()
	defer m.clientsMux.Unlock()
	m.clients[client.ID] = client
}

// RemoveClient removes a client from the model
func (m *Model) RemoveClient(clientID string) {
	m.clientsMux.Lock()
	defer m.clientsMux.Unlock()
	delete(m.clients, clientID)
}

// GetAllClients returns all connected clients
func (m *Model) GetAllClients() []*Client {
	m.clientsMux.RLock()
	defer m.clientsMux.RUnlock()

	clients := make([]*Client, 0, len(m.clients))
	for _, client := range m.clients {
		clients = append(clients, client)
	}
	return clients
}
