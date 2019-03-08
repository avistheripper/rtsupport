package main

import (
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
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	http.ListenAndServe(":8000", router)

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting single book")
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
