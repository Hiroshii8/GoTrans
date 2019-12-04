package Server

import (
	"github.com/Hiroshii8/GoTrans/Handler"
	"net/http"
)

type Server struct {
	port string
}

func (s *Server) Start() error {
	h := Handler.Handler{}
	http.HandleFunc("/", h.TranslateHandler)
	return http.ListenAndServe(s.port, nil)
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}
