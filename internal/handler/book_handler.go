package handler

import (
	"golang-rest-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	// Parse book ID from request
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Call service to fetch book by ID
	book, err := h.service.GetBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book"})
		return
	}

	// Return book as JSON response
	c.JSON(http.StatusOK, book)
}
