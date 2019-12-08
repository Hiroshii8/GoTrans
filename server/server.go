package server

import (
	"net/http"

	"github.com/Hiroshii8/GoTrans/handler"
)

// Server is struct type
type Server struct {
	port string
}

// Start ..
func (s *Server) Start() error {
	h := Handler.Handler{}
	http.HandleFunc("/", h.TranslateHandler)
	return http.ListenAndServe(s.port, nil)
}

// NewServer ..
func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}
