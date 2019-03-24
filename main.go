package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "First book", Author: "Bobby", Year: "2011"},
		Book{ID: 2, Title: "Second book", Author: "Mary", Year: "2012"},
		Book{ID: 3, Title: "Third book", Author: "Charly", Year: "2013"},
		Book{ID: 4, Title: "Fourth book", Author: "Tony", Year: "2014"},
		Book{ID: 5, Title: "Fifth book", Author: "Eliot", Year: "2016"},
	)
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	http.ListenAndServe(":8000", router)

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _,book := range books {
		if book.ID === params["id"] {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Removing book")
}
