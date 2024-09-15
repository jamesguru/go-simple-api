package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Name   string `json:"name" xml:"name" validate:"required"`
	Author string `json:"author" validate:"required"`
}

var books []Book

func AddBooks(c echo.Context) error {
	var book Book
	err := c.Bind(&book)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if book.Name == "" {
		return c.JSON(http.StatusPreconditionFailed, "name cannot be empty")
	}
	if book.Author == "" {
		return c.JSON(http.StatusPreconditionFailed, "author cannot be empty")
	}
	log.Printf("name: %s", book.Name)
	log.Printf("Author: %s", book.Author)
	books = append(books, book)
	return c.JSON(http.StatusOK, books)
}
func GetBook(c echo.Context) error {
	id := c.Param("id")
	idx, _ := strconv.ParseInt(id, 10, 64)
	return c.JSON(http.StatusOK, books[idx])
}
