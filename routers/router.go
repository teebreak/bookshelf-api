package routers

import (
	"bookshelf-api/handlers"
	middleware "bookshelf-api/middlewares"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/refresh", handlers.Refresh).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)
	api.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	api.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	api.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	api.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	api.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	api.HandleFunc("/books/index", handlers.IndexBook).Methods("POST")
	api.HandleFunc("/books/search/{query}", handlers.SearchBooks).Methods("GET")

	return router
}
