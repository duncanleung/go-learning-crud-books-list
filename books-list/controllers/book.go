package controllers

import (
	"books-list/models"
	bookRepos "books-list/repository/book"
	"books-list/utils"
	"log"

	"database/sql"
	"net/http"
)

type Controller struct {
}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		books = []models.Book{}
		bookRepo := bookRepos.BookRepository{}

		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Message = "server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}
