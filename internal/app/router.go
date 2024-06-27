package app

import (
	"go1/config"
	"go1/internal/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	cfg    config.Config
}

func NewServer(cfg config.Config) *Server {
	server := &Server{
		router: gin.Default(),
		cfg:    cfg,
	}

	// Set up your router and other configurations using cfg
	bookHandler := handler.NewBookHandler(cfg)
	server.setupRoutes(bookHandler)

	return server
}

func (s *Server) setupRoutes(bookHandler *handler.BookHandler) {
	// Implement routing logic here using s.router and bookHandler
	api := s.router.Group("/api")
	{
		api.GET("/books", bookHandler.GetBooks)
		api.GET("/books/:id", bookHandler.GetBookByID)
		api.POST("/books", bookHandler.CreateBook)
	}
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.cfg.Server.Port)
}
