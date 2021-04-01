package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID     string `json:id`
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
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get one book")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Adds one book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Updates a book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remobes a book")
}



