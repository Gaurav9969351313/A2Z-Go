package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"v3BooksRevamped/models"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

// create table books (id serial, title varchar, author varchar, year varchar);
// _ "github.com/gorilla/mux"
// _ "github.com/subosito/gotenv"
// _ "database/sql"

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	gotenv.Load()
}

var books []models.Book
var db *sql.DB

func main() {
	log.Println(os.Getenv("APP_VERSION"))

	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBooks).Methods("POST")

	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets Books")
	var book models.Book
	var books = []models.Book{}

	rows, err := db.Query("select * from books")
	logFatal(err)

	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Author, &book.Title, &book.Year)
		logFatal(err)

		books = append(books, book)
	}

	defer rows.Close()

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets Book")
	var book models.Book
	params := mux.Vars(r)
	rows := db.QueryRow("select * from books where id =$1", params["id"])
	err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year)
	logFatal(err)
	json.NewEncoder(w).Encode(book)
}

func addBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Add Book")
	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	err := db.QueryRow("insert into books(title,author,year) values($1,$2,$3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&bookID)
	logFatal(err)
	json.NewEncoder(w).Encode(bookID)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Book")
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)

	res, err := db.Exec("update books set title=$1,author=$2,year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.Id)

	rowsUpdated, err := res.RowsAffected()
	logFatal(err)
	json.NewEncoder(w).Encode(rowsUpdated)

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Book")
	params := mux.Vars(r)
	res, err := db.Exec("delete from books where id =$1", params["id"])
	logFatal(err)
	rowsDeleted, err := res.RowsAffected()
	logFatal(err)
	json.NewEncoder(w).Encode(rowsDeleted)

}
