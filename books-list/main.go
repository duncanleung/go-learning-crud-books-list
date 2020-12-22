package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"books-list/controllers"
	"books-list/driver"
	"books-list/models"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	fmt.Println("Server is running  at port 3008")
	log.Fatal(http.ListenAndServe(":3008", router))
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	row := db.QueryRow("select * from books where id=$1", params["id"])

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)
	logFatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&book)

	log.Println(book)

	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id", &book.Title, &book.Author, &book.Year, params["id"])
	logFatal(err)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Exec("delete from books where id = $1", params["id"])
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)
}
