package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ChatClient struct {
	// Connection
	conn      net.Conn
	connLock  sync.Mutex
	reader    *bufio.Reader
	connected bool

	// UI
	window     *app.Window
	th         *material.Theme
	serverIP   *widget.Editor
	username   *widget.Editor
	connectBtn *widget.Clickable
	msgInput   *widget.Editor
	sendBtn    *widget.Clickable
	messages   []string
	msgLock    sync.Mutex
}

func NewChatClient() *ChatClient {
	w := app.NewWindow(
		app.Title("Chat Client"),
		app.Size(unit.Dp(400), unit.Dp(600)),
	)

	th := material.NewTheme()

	serverIP := &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	serverIP.SetText("127.0.0.1:8080")

	username := &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}

	return &ChatClient{
		window:     w,
		th:         th,
		serverIP:   serverIP,
		username:   username,
		connectBtn: &widget.Clickable{},
		msgInput:   &widget.Editor{SingleLine: true, Submit: true},
		sendBtn:    &widget.Clickable{},
		messages:   []string{},
	}
}

func (c *ChatClient) Connect() error {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if c.connected {
		return nil
	}

	serverAddr := strings.TrimSpace(c.serverIP.Text())
	username := strings.TrimSpace(c.username.Text())

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

	c.conn = conn
	c.reader = bufio.NewReader(conn)

	_, err = c.reader.ReadString(':')
	if err != nil {
		conn.Close()
		return fmt.Errorf("failed to read welcome message: %v", err)
	}

	_, err = c.conn.Write([]byte(username + "\n"))
	if err != nil {
		conn.Close()
		return fmt.Errorf("failed to send username: %v", err)
	}

	c.connected = true

	go c.readMessages()

	return nil
}

func (c *ChatClient) Disconnect() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
		c.connected = false
	}
}

func (c *ChatClient) SendMessage(msg string) error {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if !c.connected || c.conn == nil {
		return fmt.Errorf("not connected to server")
	}

	_, err := c.conn.Write([]byte(msg + "\n"))
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	return nil
}

func (c *ChatClient) readMessages() {
	for {
		c.connLock.Lock()
		if c.conn == nil {
			c.connLock.Unlock()
			return
		}
		reader := c.reader
		c.connLock.Unlock()

		message, err := reader.ReadString('\n')
		if err != nil {
			c.AddMessage("Disconnected from server: " + err.Error())
			c.Disconnect()
			return
		}

		c.AddMessage(strings.TrimSpace(message))
	}
}

func (c *ChatClient) AddMessage(msg string) {
	c.msgLock.Lock()
	defer c.msgLock.Unlock()

	c.messages = append(c.messages, msg)
	c.window.Invalidate()
}

func (c *ChatClient) Run() error {
	var ops op.Ops

	for {
		e := <-c.window.Events()

		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			if c.connectBtn.Clicked() {
				if !c.connected {
					err := c.Connect()
					if err != nil {
						c.AddMessage("Connection error: " + err.Error())
					} else {
						c.AddMessage("Connected to server")
					}
				} else {
					c.Disconnect()
					c.AddMessage("Disconnected from server")
				}
			}

			if (c.sendBtn.Clicked() || c.msgInput.Submit) && c.connected {
				msg := strings.TrimSpace(c.msgInput.Text())
				if msg != "" {
					err := c.SendMessage(msg)
					if err != nil {
						c.AddMessage("Failed to send message: " + err.Error())
					}
					c.msgInput.SetText("")
				}
				c.msgInput.Submit = false
			}

			if !c.connected {
				c.renderLoginScreen(gtx)
			} else {
				c.renderChatScreen(gtx)
			}

			e.Frame(gtx.Ops)
		}
	}
}

func (c *ChatClient) renderLoginScreen(gtx layout.Context) {
	layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceStart,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top:    unit.Dp(20),
				Bottom: unit.Dp(20),
				Left:   unit.Dp(20),
				Right:  unit.Dp(20),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Bottom: unit.Dp(20)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							title := material.H4(c.th, "Chat Client")
							title.Alignment = text.Middle
							return title.Layout(gtx)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Bottom: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									label := material.Body1(c.th, "Server IP: ")
									return label.Layout(gtx)
								}),
								layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
									return material.Editor(c.th, c.serverIP, "").Layout(gtx)
								}),
							)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Bottom: unit.Dp(20)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									label := material.Body1(c.th, "User Name: ")
									return label.Layout(gtx)
								}),
								layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
									return material.Editor(c.th, c.username, "").Layout(gtx)
								}),
							)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(c.th, c.connectBtn, "Connect")
						return layout.Inset{Left: unit.Dp(100), Right: unit.Dp(100)}.Layout(gtx, btn.Layout)
					}),
				)
			})
		}),
	)
}

func (c *ChatClient) renderChatScreen(gtx layout.Context) {
	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top:    unit.Dp(10),
				Bottom: unit.Dp(10),
				Left:   unit.Dp(10),
				Right:  unit.Dp(10),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// Chat history area with border
				return widget.Border{
					Color: color.NRGBA{A: 200},
					Width: unit.Dp(1),
				}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    unit.Dp(5),
						Bottom: unit.Dp(5),
						Left:   unit.Dp(5),
						Right:  unit.Dp(5),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						c.msgLock.Lock()
						messages := c.messages
						c.msgLock.Unlock()

						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								messageList := layout.List{
									Axis: layout.Vertical,
								}
								return messageList.Layout(gtx, len(messages), func(gtx layout.Context, i int) layout.Dimensions {
									return layout.Inset{
										Top:    unit.Dp(2),
										Bottom: unit.Dp(2),
										Left:   unit.Dp(5),
										Right:  unit.Dp(5),
									}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										msg := material.Body1(c.th, messages[i])
										return msg.Layout(gtx)
									})
								})
							}),
						)
					})
				})
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Bottom: unit.Dp(10),
				Left:   unit.Dp(10),
				Right:  unit.Dp(10),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Horizontal,
					Alignment: layout.Middle,
				}.Layout(gtx,
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(c.th, c.msgInput, "Type a message...").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(c.th, c.sendBtn, "Send")
							return btn.Layout(gtx)
						})
					}),
				)
			})
		}),
	)
}

func main() {
	go func() {
		client := NewChatClient()
		if err := client.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
