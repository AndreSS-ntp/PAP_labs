package view

import (
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab2/internal/model"
)

// View handles the presentation logic
type View struct{}

// NewView creates a new view instance
func NewView() *View {
	return &View{}
}

// BroadcastMessage sends a message to all specified clients
func (v *View) BroadcastMessage(content string, sender string, clients []*model.Client) {
	var formattedMsg string

	if sender == "" {
		// System message
		formattedMsg = fmt.Sprintf("[СИСТЕМА] %s", content)
	} else {
		// User message
		formattedMsg = fmt.Sprintf("[%s] %s\n", sender, content)
	}

	// Send the message to all clients
	for _, client := range clients {
		_, err := client.Conn.Write([]byte(formattedMsg))
		if err != nil {
			// Silently ignore errors, they'll be handled when the client disconnects
			continue
		}
	}
}
