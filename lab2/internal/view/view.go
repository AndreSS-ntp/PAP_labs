package view

import (
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab2/internal/model"
)

type View struct{}

func NewView() *View {
	return &View{}
}

func (v *View) BroadcastMessage(content string, sender string, clients []*model.Client) {
	var formattedMsg string

	if sender == "" {
		formattedMsg = fmt.Sprintf("[СИСТЕМА] %s", content)
	} else {
		formattedMsg = fmt.Sprintf("[%s] %s\n", sender, content)
	}
	
	for _, client := range clients {
		_, err := client.Conn.Write([]byte(formattedMsg))
		if err != nil {
			// Silently ignore errors, they'll be handled when the client disconnects
			continue
		}
	}
}
