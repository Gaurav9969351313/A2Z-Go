package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books,
		Book{Id: 0, Title: "A", Author: "G", Year: "2017"},
		Book{Id: 1, Title: "B", Author: "K", Year: "2018"},
		Book{Id: 2, Title: "C", Author: "R", Year: "2019"},
		Book{Id: 3, Title: "D", Author: "Y", Year: "2020"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBooks).Methods("POST")

	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets Books")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets Book")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.Id == id {
			json.NewEncoder(w).Encode(&book)
		}
	}
	// log.Println(params) //map[id:11]
	// error (mismatched types int and string)
	// reflect.TypeOf(params["id"]) /string
}

func addBooks(w http.ResponseWriter, r *http.Request) {

	log.Println("Add Book")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Book")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.Id == book.Id {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete Book")

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for i, book := range books {
		if book.Id == id {
			books = append(books[:i], books[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(books)
}
