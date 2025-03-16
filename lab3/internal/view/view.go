package view

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type View struct {
	// Theme
	th *material.Theme

	// UI elements
	serverIP   *widget.Editor
	username   *widget.Editor
	connectBtn *widget.Clickable
	msgInput   *widget.Editor
	sendBtn    *widget.Clickable
}

func NewView(theme *material.Theme) *View {
	serverIP := &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	serverIP.SetText("127.0.0.1:8080")

	username := &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}

	return &View{
		th:         theme,
		serverIP:   serverIP,
		username:   username,
		connectBtn: &widget.Clickable{},
		msgInput:   &widget.Editor{SingleLine: true, Submit: true},
		sendBtn:    &widget.Clickable{},
	}
}

func (v *View) GetServerIP() string {
	return v.serverIP.Text()
}

func (v *View) GetUsername() string {
	return v.username.Text()
}

func (v *View) GetMessageInput() string {
	return v.msgInput.Text()
}

func (v *View) ClearMessageInput() {
	v.msgInput.SetText("")
}

func (v *View) IsConnectClicked() bool {
	return v.connectBtn.Clicked()
}

func (v *View) IsSendClicked() bool {
	return v.sendBtn.Clicked() || v.msgInput.Submit
}

func (v *View) ResetSubmit() {
	v.msgInput.Submit = false
}

func (v *View) RenderLoginScreen(gtx layout.Context) layout.Dimensions {
	return layout.Flex{
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
							title := material.H4(v.th, "Chat Client")
							title.Alignment = text.Middle
							return title.Layout(gtx)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Bottom: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									label := material.Body1(v.th, "Server IP: ")
									// Make label bold
									return label.Layout(gtx)
								}),
								layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
									return material.Editor(v.th, v.serverIP, "").Layout(gtx)
								}),
							)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Bottom: unit.Dp(20)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									label := material.Body1(v.th, "User Name: ")
									return label.Layout(gtx)
								}),
								layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
									return material.Editor(v.th, v.username, "").Layout(gtx)
								}),
							)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(v.th, v.connectBtn, "Connect")
						return layout.Inset{Left: unit.Dp(100), Right: unit.Dp(100)}.Layout(gtx, btn.Layout)
					}),
				)
			})
		}),
	)
}

func (v *View) RenderChatScreen(gtx layout.Context, messages []string) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top:    unit.Dp(10),
				Bottom: unit.Dp(10),
				Left:   unit.Dp(10),
				Right:  unit.Dp(10),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
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
										msg := material.Body1(v.th, messages[i])
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
						return material.Editor(v.th, v.msgInput, "Type a message...").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(v.th, v.sendBtn, "Send")
							return btn.Layout(gtx)
						})
					}),
				)
			})
		}),
	)
}
