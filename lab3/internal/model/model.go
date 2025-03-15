package model

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Model struct {
	// Connection
	conn      net.Conn
	connLock  sync.Mutex
	reader    *bufio.Reader
	connected bool

	// Data
	messages []string
	msgLock  sync.Mutex
	serverIP string
	username string
}

func NewModel() *Model {
	return &Model{
		messages: []string{},
		serverIP: "127.0.0.1:8080",
		username: "",
	}
}

func (m *Model) Connect(serverAddr, username string) error {
	m.connLock.Lock()
	defer m.connLock.Unlock()

	if m.connected {
		return nil
	}

	if serverAddr == "" {
		return fmt.Errorf("server address cannot be empty")
	}
	if username == "" {
		return fmt.Errorf("username cannot be empty")
	}

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}

	m.conn = conn
	m.reader = bufio.NewReader(conn)
	m.serverIP = serverAddr
	m.username = username

	_, err = m.reader.ReadString(':')
	if err != nil {
		conn.Close()
		return fmt.Errorf("failed to read welcome message: %v", err)
	}

	_, err = m.conn.Write([]byte(username + "\n"))
	if err != nil {
		conn.Close()
		return fmt.Errorf("failed to send username: %v", err)
	}

	m.connected = true
	return nil
}

func (m *Model) Disconnect() {
	m.connLock.Lock()
	defer m.connLock.Unlock()

	if m.conn != nil {
		m.conn.Close()
		m.conn = nil
		m.connected = false
	}
}

func (m *Model) SendMessage(msg string) error {
	m.connLock.Lock()
	defer m.connLock.Unlock()

	if !m.connected || m.conn == nil {
		return fmt.Errorf("not connected to server")
	}

	// Send message
	_, err := m.conn.Write([]byte(msg + "\n"))
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	return nil
}

func (m *Model) AddMessage(msg string) {
	m.msgLock.Lock()
	defer m.msgLock.Unlock()

	m.messages = append(m.messages, msg)
}

func (m *Model) GetMessages() []string {
	m.msgLock.Lock()
	defer m.msgLock.Unlock()

	messagesCopy := make([]string, len(m.messages))
	copy(messagesCopy, m.messages)

	return messagesCopy
}

func (m *Model) IsConnected() bool {
	m.connLock.Lock()
	defer m.connLock.Unlock()

	return m.connected
}

func (m *Model) GetReader() *bufio.Reader {
	m.connLock.Lock()
	defer m.connLock.Unlock()

	return m.reader
}

func (m *Model) GetServerIP() string {
	return m.serverIP
}

func (m *Model) SetServerIP(ip string) {
	m.serverIP = ip
}

func (m *Model) GetUsername() string {
	return m.username
}

func (m *Model) SetUsername(username string) {
	m.username = username
}
