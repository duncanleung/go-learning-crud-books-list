package bookRepository

import (
	"books-list/models"
	"database/sql"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {

	rows, err := db.Query("select * from books")
	if err != nil {
		return []models.Book{}, err
	}

	// defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

		books = append(books, book)
	}
	if err != nil {
		return []models.Book{}, err
	}

	return books, err
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) (models.Book, error) {

	row := db.QueryRow("select * from books where id=$1", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	return book, err
}
