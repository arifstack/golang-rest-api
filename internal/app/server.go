package app

import (
	"fmt"
	"net/http"
)

type Server struct {
	port   int
	router http.Handler
}

func NewServer(port int, router http.Handler) *Server {
	return &Server{
		port:   port,
		router: router,
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.port)
	return http.ListenAndServe(addr, s.router)
}
