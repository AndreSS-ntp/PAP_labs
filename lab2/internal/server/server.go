package server

import (
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab2/internal/controller"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab2/internal/model"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab2/internal/view"
	"net"
)

// Server represents the chat server
type Server struct {
	address    string
	listener   net.Listener
	model      *model.Model
	view       *view.View
	controller *controller.Controller
}

// NewServer creates a new server instance
func NewServer(address string) *Server {
	model := model.NewModel()
	view := view.NewView()
	controller := controller.NewController(model, view)

	return &Server{
		address:    address,
		model:      model,
		view:       view,
		controller: controller,
	}
}

// Start starts the server and begins accepting connections
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to start listener: %v", err)
	}
	s.listener = listener

	fmt.Printf("Сервер запущен на %s\n", s.address)
	fmt.Println("Ожидание подключений...")

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Printf("Ошибка при принятии соединения: %v\n", err)
			continue
		}

		// Handle each connection in a separate goroutine
		go s.controller.HandleConnection(conn)
	}
}

// Stop stops the server
func (s *Server) Stop() error {
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}
