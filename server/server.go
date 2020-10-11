package server

import (
	"fmt"
	"net/http"

	"github.com/Hiroshii8/GoTrans/handler"
)

// Server is struct type
type Server struct {
	port string
}

// Start server action
func (s *Server) Start() error {
	h := handler.Handler{}
	http.HandleFunc("/", h.TranslateHandler)
	fmt.Println("Starting Server at", s.port)
	return http.ListenAndServe(s.port, nil)
}

// NewServer will initialize Server Object
func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}
