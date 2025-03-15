package controller

import (
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab3/internal/model"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab3/internal/view"
	"strings"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
)

type Controller struct {
	window *app.Window
	model  *model.Model
	view   *view.View
}

func NewController(window *app.Window, model *model.Model, view *view.View) *Controller {
	return &Controller{
		window: window,
		model:  model,
		view:   view,
	}
}

func (c *Controller) Run() error {
	var ops op.Ops

	go c.readMessages()

	for {
		e := <-c.window.Events()

		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			if c.view.IsConnectClicked() {
				if !c.model.IsConnected() {
					serverAddr := strings.TrimSpace(c.view.GetServerIP())
					username := strings.TrimSpace(c.view.GetUsername())

					err := c.model.Connect(serverAddr, username)
					if err != nil {
						c.model.AddMessage("Connection error: " + err.Error())
					} else {
						c.model.AddMessage("Connected to server")
					}
				} else {
					c.model.Disconnect()
					c.model.AddMessage("Disconnected from server")
				}
				c.window.Invalidate()
			}

			if c.view.IsSendClicked() && c.model.IsConnected() {
				msg := strings.TrimSpace(c.view.GetMessageInput())
				if msg != "" {
					err := c.model.SendMessage(msg)
					if err != nil {
						c.model.AddMessage("Failed to send message: " + err.Error())
					}
					c.view.ClearMessageInput()
				}
				c.view.ResetSubmit()
				c.window.Invalidate()
			}

			// Render UI
			if !c.model.IsConnected() {
				c.view.RenderLoginScreen(gtx)
			} else {
				c.view.RenderChatScreen(gtx, c.model.GetMessages())
			}

			e.Frame(gtx.Ops)
		}
	}
}

func (c *Controller) readMessages() {
	for {
		if !c.model.IsConnected() {
			// Sleep a bit to avoid busy waiting
			// This will be reset when we connect again
			continue
		}

		reader := c.model.GetReader()
		if reader == nil {
			continue
		}

		message, err := reader.ReadString('\n')
		if err != nil {
			c.model.AddMessage("Disconnected from server: " + err.Error())
			c.model.Disconnect()
			c.window.Invalidate()
			continue
		}

		c.model.AddMessage(strings.TrimSpace(message))
		c.window.Invalidate()
	}
}
