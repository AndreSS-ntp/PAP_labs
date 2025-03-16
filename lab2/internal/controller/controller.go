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

type Controller struct {
	model *model.Model
	view  *view.View
}

func NewController(model *model.Model, view *view.View) *Controller {
	return &Controller{
		model: model,
		view:  view,
	}
}

func (c *Controller) HandleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Введите ваше имя: "))
	scanner := bufio.NewScanner(conn)

	if !scanner.Scan() {
		return
	}

	username := strings.TrimSpace(scanner.Text())
	if username == "" {
		username = "Аноним"
	}

	clientID := fmt.Sprintf("%s-%d", conn.RemoteAddr().String(), time.Now().UnixNano())
	client := &model.Client{
		ID:       clientID,
		Conn:     conn,
		Username: username,
	}

	c.model.AddClient(client)

	joinMsg := fmt.Sprintf("%s присоединился к чату\n", username)
	c.view.BroadcastMessage(joinMsg, "", c.model.GetAllClients())

	welcomeMsg := "Добро пожаловать в чат! Введите сообщение и нажмите кнопку для отправки.\n"
	conn.Write([]byte(welcomeMsg))

	for scanner.Scan() {
		text := scanner.Text()

		if strings.ToLower(text) == "/exit" {
			break
		}

		message := &model.Message{
			Content:   text,
			Sender:    username,
			Timestamp: time.Now().Format("15:04:05"),
		}

		c.view.BroadcastMessage(message.Content, message.Sender, c.model.GetAllClients())
	}

	c.model.RemoveClient(clientID)

	leaveMsg := fmt.Sprintf("%s покинул чат\n", username)
	c.view.BroadcastMessage(leaveMsg, "", c.model.GetAllClients())
}
