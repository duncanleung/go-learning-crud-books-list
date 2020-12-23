package controllers

import (
	"books-list/models"
	bookRepos "books-list/repository/book"
	"books-list/utils"
	"encoding/json"
	"log"
	"strconv"

	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
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
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		params := mux.Vars(r)

		books = []models.Book{}
		bookRepo := bookRepos.BookRepository{}

		id, _ := strconv.Atoi(params["id"])

		book, err := bookRepo.GetBook(db, book, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not Found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		var bookID int

		json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Please enter Author, Title, and Year"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepos.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book, bookID)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Please enter Author, Title, and Year"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepos.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book, id)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsUpdated)
	}
}
