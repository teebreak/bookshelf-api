package routers

import (
	"bookshelf-api/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/api/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", handlers.DeleteBook).Methods("DELETE")
	return router
}
