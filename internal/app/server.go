package app

import (
	"go1/internal/db"
	"go1/internal/handler"
	"go1/internal/repository"
	"go1/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	dbConn := db.Connect()
	bookRepo := repository.NewMySQLBookRepository(dbConn)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()
	router.GET("/books", bookHandler.GetBooks)
	router.GET("/books/:id", bookHandler.GetBookByID)
	router.POST("/books", bookHandler.CreateBook)

	return &Server{router: router}
}

func (s *Server) Run() error {
	return s.router.Run(":8080")
}
