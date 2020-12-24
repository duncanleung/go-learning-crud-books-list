# go-learning-crud-books-list

Learning Go through a udemy tutorial:

[Golang: Intro to REST APIs with Go programming lang (Golang)](https://www.udemy.com/course/build-a-restful-api-with-golang-go-programming-language/)
https://github.com/codixir/books-list-with-postgres

## SQL Setup

Uses https://www.elephantsql.com/

Create `.env`

```
POSTGRESQL_URL = "postgres://{{DB_NAME}}:{{DB_PASSWORD}}@suleiman.db.elephantsql.com:5432/{{DB_NAME}}"
```

Database connection is created in:

`/driver/driver.go`

## Go Server

Routes of the Go server are:

```go
router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
router.HandleFunc("/books/{id}", controller.UpdateBook(db)).Methods("PUT")
router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")
```

SQL logic for each handler is extracted out to `/repository/book/book_psql.go`

```go
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) (models.Book, error) {
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book, bookID int) (int, error) {
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book, id int) (int64, error) {
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) (int64, error) {
}
```
