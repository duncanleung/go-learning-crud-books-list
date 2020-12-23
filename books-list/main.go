package main

import (
	"database/sql"
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
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books/{id}", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is running  at port 3008")
	log.Fatal(http.ListenAndServe(":3008", router))
}
