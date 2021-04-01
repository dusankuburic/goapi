package main

import (
	"book-list/driver"
	"book-list/models"
	"book-list/controllers"
	"net/http"

	"database/sql"
	_ "database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	_ "github.com/subosito/gotenv"
	"log"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var books []models.Book
var db *sql.DB

func init() {
	err := gotenv.Load()
	logFatal(err)
}

func main() {

	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.GetBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
