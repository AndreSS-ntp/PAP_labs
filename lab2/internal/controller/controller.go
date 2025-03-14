package controller

import (
	"bufio"
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab2/internal/model"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab2/internal/view"
	"net"
	"strings"
	"time"
)

// Controller handles the business logic of the chat server
type Controller struct {
	model *model.Model
	view  *view.View
}

// NewController creates a new controller instance
func NewController(model *model.Model, view *view.View) *Controller {
	return &Controller{
		model: model,
		view:  view,
	}
}

// HandleConnection handles a new client connection
func (c *Controller) HandleConnection(conn net.Conn) {
	defer conn.Close()

	// Ask for username
	conn.Write([]byte("Введите ваше имя: "))
	scanner := bufio.NewScanner(conn)

	// Wait for username
	if !scanner.Scan() {
		return
	}

	username := strings.TrimSpace(scanner.Text())
	if username == "" {
		username = "Аноним"
	}

	// Create a new client
	clientID := fmt.Sprintf("%s-%d", conn.RemoteAddr().String(), time.Now().UnixNano())
	client := &model.Client{
		ID:       clientID,
		Conn:     conn,
		Username: username,
	}

	// Add client to the model
	c.model.AddClient(client)

	// Broadcast join message
	joinMsg := fmt.Sprintf("%s присоединился к чату\n", username)
	c.view.BroadcastMessage(joinMsg, "", c.model.GetAllClients())

	// Welcome message
	welcomeMsg := "Добро пожаловать в чат! Введите сообщение и нажмите Enter для отправки.\n"
	conn.Write([]byte(welcomeMsg))

	// Handle client messages
	for scanner.Scan() {
		text := scanner.Text()

		// Check if client wants to exit
		if strings.ToLower(text) == "/exit" {
			break
		}

		// Create message
		message := &model.Message{
			Content:   text,
			Sender:    username,
			Timestamp: time.Now().Format("15:04:05"),
		}

		// Broadcast message to all clients
		c.view.BroadcastMessage(message.Content, message.Sender, c.model.GetAllClients())
	}

	// Remove client when they disconnect
	c.model.RemoveClient(clientID)

	// Broadcast leave message
	leaveMsg := fmt.Sprintf("%s покинул чат\n", username)
	c.view.BroadcastMessage(leaveMsg, "", c.model.GetAllClients())
}
