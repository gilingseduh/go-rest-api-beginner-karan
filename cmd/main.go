package main

import (
	"github.com/gorilla/mux"
	"go-rest-api-beginner-karan/pkg/db"
	"go-rest-api-beginner-karan/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	database := db.Init()
	handler := handlers.New(database)
	router := mux.NewRouter()

	router.HandleFunc("/books", handler.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handler.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", handler.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", handler.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", handler.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running...")
	http.ListenAndServe(":4000", router)
}
